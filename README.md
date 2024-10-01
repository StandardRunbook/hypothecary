# hypothecary
Standardized runbook that helps engineers automate runbook tasks in your cluster before engineers are paged

TLDR: Easy-to-install k8s job to run system debugging hypotheses

All too often, debugging gnarly microservices issues involves several steps:
1. get production approval
2. ssh into the bastion
3. ssh into the container
4. run curl command or script (hopefully you don't misspell the command in the middle of the night)

You might want to test if you can `ping a dependency endpoint` from that individual host. Or you want to `renew a certificate` or `retokenike the puppet key`.

Ultimately, debugging is about validating hypotheses and these commands are effectively testing some aspect of the system to narrow down the issue.

Runbooks (or playbooks if you're at Google) are filled with these commands, but I find the process of following runbooks cumbersome, error-prone and frankly annoying when I get paged at 4 in the morning. Sometimes, the runbooks are wrong, missing context, or simply not updated and I simply can't rely on them as a source of truth. Fixing this is a culture/process problem but stocks vest faster than culture changes.

So I propose standardizing runbooks with pre-configured scripts instead of open form text. 

I think container logs are a great source of truth. So why not run these commands directly in the container to run validating scripts with my dependencies, check the expiry of the certificate, etc. before I even get paged.
This way, when I actually get paged, I'm not sitting around waiting for an approval or parsing through logs/metrics/traces in the middle of night and I can figure out the source of the issue faster.

In that spirit, I offer an alternative: why not have a pre-configured set of scripts that run when the readiness/liveness/startup probes fail and output the responses to your preferred logging sink, e.g. Splunk.

For example, running a command to test if you can connect to AWS tells you:

- cluster + pod + container is still working
- network policy for outbound requests are set correctly
- NAT gateway/outbound proxy is working normally
- AWS is not down
- more stuff in between

Running another command to connect to Digicert shares the above path and only differs at the leaf. We can show by overlapping the 

As Pranav Mistry says: `we as humans are not interested in technology [or processes], we're interested in information.`
In that mode of thinking, I have a service that will automagically root cause incidents from these hypotheses with a bayesian network called BayesicInsight. This insight service is not open source currently.

# Design

This library is designed with a plugin system in mind (called hypotheses). A hypothesis is an abstraction that can be:

- pre-built plugin, e.g. `AWS-STS`
- shell command, e.g. `curl -X autorootcause.com`
- script, e.g. `sh renew_certificate.sh`

This format is useful for two main reasons:

- standardized hypotheses can be shared between teams within a company and between companies
- they can be chained to create conditionals and pipelines, letting you auto-diagnose and sometimes even auto-remediate directly in the cluster

After the hypotheses are run, a report is generated at the end.

Currently, connectors to Datadog, Splunk, etc. are not built and have to instrumented by teams on their own but the goal is to send the results to Datadog, Splunk, Grafana or whatever log/metric storage you use.

### Pre-built Plugins

With the number of folks building managed services, there's a lot of duplication in how the industry validates whether a service is up. So if you'd like to use one of the pre-built plugins, add one of the plugins to your config:
```yaml
plugins:
  - plugin:
      name: AWS-STS
      type: prebuilt
      conditionals:
        - response: http_2xx
          child_name: plugin-do-something
        - response: not_http_2xx
          child_name: None
      parallel: true
      trigger:
        trigger_type:
          - on_liveness_failure
          - periodic
        trigger_interval: 5
        trigger_timeout: 60s
```

### Shell Command

```yaml
plugins:
  - plugin:
      name: call-openai
      type: shell
      parallel: true          
      shell_command: curl -X GET https://api.openai.com/v1/engines
      conditionals: 
        - response: http_2xx 
          child_name: plugin-do-something
        - response: not_http_2xx
          child_name: skip-to-end
      trigger:
        trigger_type:
          - on_liveness_failure
          - periodic
        trigger_interval: 5
        trigger_timeout: 60s
```

### Shell Scripts

It's assumed that the file has the right permissions before this runs.

```yaml
plugins:
  - plugin:
      name: call-openai
      type: shell
      parallel: true
      shell_command: sh script.sh
      conditionals:
        - response: http_2xx
          child_name: plugin-do-something
        - response: not_http_2xx
          child_name: plugin-do-something-else
      trigger:
        trigger_type:
          - on_liveness_failure
          - periodic
        trigger_interval: 5
        trigger_timeout: 60s
```

## Plugins

Add your plugin to this list once you merge it into the plugin folder.

- aws-get-caller-identity: calls aws get caller identity and checks if the response code is 200
- ls-file-exists: checks to see if a file was placed successfully
- cert-expiry-check: checks the expiry of a certificate
- file-permissions-check: matches the user given permissions to a current file
