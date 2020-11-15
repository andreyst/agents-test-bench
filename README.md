# agents-test-bench
Load &amp; failure injection test bench of various logging/monitoring agents.

# Prepare

| **NB:** Following instructions install a number of packages that you might not need or want on a system on a permanent basis. Provisioning a node, virtual machine or container specifically for test bench is recommended. |
| --- |

1. Get [Salt](https://docs.saltstack.com/en/latest/topics/installation/index.html) on target system where you want to deploy agents test bench.
2. Clone this repo.
3. Configure test bench working directory by copying `pillar/local_config.sls.example` to `pillar/local_config.sls` and changing directory to desired. This directory will contain all configs, binaries and most importantly data, so ensure that directory is on partition which has enough disk space. Needed disk space depends on test duration, tens of GBs recommended.
4. Apply salt state with `salt-call --log-level debug --local --file-root ./salt --pillar-root ./pillar state.highstate`. This should be executed from cloned repo root. This operation might take a while to install or build everything.
5. Applying salt state should finish without any errors. If there are any failed states, please report it via Github issues.

# Troubleshooting

## Running highstate fails with `Specified SLS 'local_config' in environment 'base' is not available on the salt master`.

You need to create local_config.sls. See ["Prepare"](#prepare) section for instructions on how to do that.

## When applying salt state, some states fail

You can either report it via Guthub issues or try to investigate yourself. Start with the error message, find what failed and what state should do.