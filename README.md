# hypothecary
Automate runbook tasks in your cluster before you're paged

TLDR: Easy-to-install k8s job to run system debugging hypotheses

All too often, debugging gnarly microservices issues involves several steps:
1. get production approval
2. ssh into the bastion
3. ssh into the container
4. run curl command or script (hopefully you don't misspell the command in the middle of the night)

You might want to test if you can `ping a dependency endpoint` from that individual host. Or you want to `renew a certificate` or `retokenike the puppet key`.

Ultimately, debugging is about validating hypotheses and these commands are effectively testing some aspect of the system to narrow down the issue.

Runbooks (or playbooks if you're at Google) are filled with these commands, but I find the process of following runbooks cumbersome, error-prone and frankly annoying when I get paged at 4 in the morning. Sometimes, the runbooks are wrong, missing context, or simply not updated and I simply can't rely on them as a source of truth. Fixing this is a culture/process problem but stocks vest faster than culture changes.

However, I think container logs are a great source of truth. So why not run these commands directly in the container to run validating scripts with my dependencies, check the expiry of the certificate, etc. before I even get paged.
This way, when I actually get paged, I'm not sitting around waiting for an approval or parsing through logs/metrics/traces in the middle of night and I can figure out the source of the issue faster.

In that spirit, I offer an alternative: why not have a pre-configured set of scripts that run when the readiness/liveness/startup probes break and output the responses to your preferred logging sink, e.g. Splunk.

For example, running a command to test if you can connect to AWS tells you:

- cluster is still working
- network policy for outbound requests are set correctly
- NAT gateway/outbound proxy is working normally
- AWS is not down
- more stuff in between

As Pranav Mistry says: `we as humans are not interested in technology [or processes], we're interested in information.`
In that mode of thinking, I have a shameless plug: reach out if you'd like to go a step forward to use these hypotheses to automagically root cause incidents and create dashboards that combine fragmented information from metrics, logs and traces' backends.

# Design

This library is designed with a plugin system in mind. A plugin is an abstraction that can be:

- pre-built plugin, e.g. `AWS-STS`
- shell command, e.g. `curl -X autorootcause.com`
- script, e.g. `sh renew_certificate.sh`

This format is useful for two main reasons:

- standardized plugins can be shared between teams within a company and between companies
- they can be chained to create conditionals and pipelines, letting you auto-diagnose and sometimes even auto-remediate directly in the cluster

After the plugins are run, a report is generated at the end. Currently, these connectors are not built and have to instrumented by teams on their own but the goal is to send the results to Datadog, Splunk, Grafana or whatever log/metric storage you use. If you don't have a log storage, reach out, and we can set up a dashboard for you.

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
