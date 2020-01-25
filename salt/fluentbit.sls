flex:
  pkg.installed

bison:
  pkg.installed

build-essential:
  pkg.installed

zlib1g-dev:
  pkg.installed

wget:
  pkg.installed

build-fluentbit-script:
  file.managed:
    - name: {{ pillar['test_bench_directory'] }}/build/fluentbit/build-fluentbit.sh
    - source: salt://fluentbit/build-fluentbit.sh
    - makedirs: True
    - mode: 755
    - template: jinja

build-fluentbit:
  cmd.run:
    - name: {{ pillar['test_bench_directory'] }}/build/fluentbit/build-fluentbit.sh
    - unless: test -L "{{ pillar['test_bench_directory'] }}/bin/fluent-bit"
    - require:
      - pkg: wget
      - pkg: flex
      - pkg: bison
      - pkg: build-essential
      - pkg: zlib1g-dev
      - file: build-fluentbit-script

fluentbit-config:
  file.managed:
    - name: {{ pillar['test_bench_directory'] }}/config/fluentbit/fluentbit.conf
    - source: salt://fluentbit/fluentbit.conf
    - makedirs: True
    - template: jinja

