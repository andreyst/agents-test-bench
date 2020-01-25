include:
  - test_bench_dirs
  - pip
  - fluentbit

pip_supervisor:
  pip.installed:
    - name: supervisor
    - require:
      - sls: pip

supervisor_config:
  file.managed:
    - name: {{ pillar['test_bench_directory'] }}/config/supervisor/supervisord.conf
    - source: salt://supervisor/supervisord.conf
    - template: jinja
    - makedirs: True
    - require:
      - pip: pip_supervisor
      - sls: test_bench_dirs

{{ pillar['test_bench_directory'] }}/data/logs/supervisor:
  file.directory:
    - makedirs: True
    - require:
      - pip: pip_supervisor
      - sls: test_bench_dirs

supervisor_running:
  cmd.run:
    - name: supervisord -c "{{ pillar['test_bench_directory'] }}/config/supervisor/supervisord.conf"
    - unless: test -f "{{ pillar['test_bench_directory'] }}/run/supervisord.pid" && kill -0 $(cat "{{ pillar['test_bench_directory'] }}/run/supervisord.pid")
    - require:
      - pip: pip_supervisor
      - file: supervisor_config
      - sls: fluentbit
