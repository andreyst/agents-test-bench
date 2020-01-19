## Prepare
```
#!/bin/bash
set -xe

apt-get update
apt-get install nano wget curl -y
apt-get install flex bison build-essential zlib1g-dev -y
cd ~
wget https://fluentbit.io/releases/1.3/fluent-bit-1.3.6.tar.gz
tar xvf fluent-bit-1.3.6.tar.gz
curl -sSL https://cmake.org/files/v3.5/cmake-3.5.2-Linux-x86_64.tar.gz | sudo tar -xzC /opt
cd fluent-bit-1.3.6/build
/opt/cmake-3.5.2-Linux-x86_64/bin/cmake -DFLB_OUT_KAFKA=On -DWITH_ZLIB=1 ../
make

cd ..
cat <<EOF > test.conf
[INPUT]
    Name  cpu

[INPUT]
    Name  tcp
    Port  9099

[OUTPUT]
    Name        kafka
    Match       *
    Brokers     localhost:9092
    Topics      test
    rdkafka.compression.codec gzip

[OUTPUT]
    Name stdout
    Match *
EOF

./build/bin/fluent-bit -c test.conf
```
