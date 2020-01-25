{{ pillar['test_bench_directory'] }}:
  file.directory:
    - makedirs: True

{{ pillar['test_bench_directory'] }}/build:
  file.directory:
    - makedirs: True

{{ pillar['test_bench_directory'] }}/bin:
  file.directory:
    - makedirs: True

{{ pillar['test_bench_directory'] }}/config:
  file.directory:
    - makedirs: True

{{ pillar['test_bench_directory'] }}/run:
  file.directory:
    - makedirs: True

{{ pillar['test_bench_directory'] }}/data:
  file.directory:
    - makedirs: True

{{ pillar['test_bench_directory'] }}/data/logs:
  file.directory:
    - makedirs: True
