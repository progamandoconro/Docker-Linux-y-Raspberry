#33333333333333333333333333333333333333333333333333333333333333333333333333333333333
# KIND: Si tienes Docker y Go. Simplemente ejutar:

GO111MODULE="on" go get sigs.k8s.io/kind@v0.5.1 && kind create cluster
 export PATH=$PATH:$(go env GOPATH)/bin
 
 kind create cluster
export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
kubectl cluster-info
kubectl cluster-info dump

#####################################################################################3
# Si tienes solo Go
sudo -i
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/kubernetes/kubernetes
cd kubernetes
make

##################################################3

# Kubernetes for Docker users.
#https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/

############ BINARIOS
curl -Lo ./kind https://github.com/kubernetes-sigs/kind/releases/download/v0.5.1/kind-$(uname)-amd64
chmod +x ./kind
mv ./kind /some-dir-in-your-PATH/kind
