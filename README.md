Run:

    make build_darwin
    
Example: 

    errorToSigkill kubectl port-forward --address 0.0.0.0 service/crm 8080:8080
    
Systemd:

    [Unit]
    Description=Kubectl port forward for service crm
    After=network.target
    
    [Service]
    Type=simple
    ExecStart=/usr/local/bin/errorToSigkill_linux /usr/bin/kubectl --kubeconfig=/home/username/.kube/config -n default port-forward --address 0.0.0.0 service/crm 80:8080
    Restart=always
    
    [Install]
    WantedBy=default.target        