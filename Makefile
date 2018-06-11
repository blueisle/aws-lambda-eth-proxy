build:
	dep ensure
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/eth-proxy \
			./
	mv bin/eth-proxy-linux-amd64 bin/eth-proxy
