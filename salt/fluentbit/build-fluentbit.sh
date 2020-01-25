#!/bin/bash
set -xe

cd "{{ pillar['test_bench_directory'] }}/build/fluentbit/"

wget --no-clobber https://fluentbit.io/releases/1.3/fluent-bit-1.3.6.tar.gz
tar -xf fluent-bit-1.3.6.tar.gz -C "{{ pillar['test_bench_directory'] }}/build/fluentbit/"

wget --no-clobber https://cmake.org/files/v3.5/cmake-3.5.2-Linux-x86_64.tar.gz
tar -xf cmake-3.5.2-Linux-x86_64.tar.gz -C"{{ pillar['test_bench_directory'] }}/build/fluentbit/"

cd "{{ pillar['test_bench_directory'] }}/build/fluentbit/fluent-bit-1.3.6/build"
"{{ pillar['test_bench_directory'] }}/build/fluentbit/cmake-3.5.2-Linux-x86_64/bin/cmake" -DFLB_OUT_KAFKA=On -DWITH_ZLIB=1 ../
make

ln -s "{{ pillar['test_bench_directory'] }}/build/fluentbit/fluent-bit-1.3.6/build/bin/fluent-bit" "{{ pillar['test_bench_directory'] }}/bin/fluent-bit"
