
# Docker
```bash
# build
docker build -t blug:dev-1.0 .

# run
docker run -it --rm -p 8000:8000 --name blug -v ./template:/data/conf blug:dev-1.0
```

# in your vps
```bash
mv deploy.sh.backend deploy.sh
chmod +x deploy.sh
sudo apt install sshpass
#change ***** to your info
./deploy.sh
```


