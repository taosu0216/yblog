
# Use Docker and running on local
```bash
mv template/config.yaml.template template/config.yaml
mv cmd/ent/migration.go.backup cmd/ent/migration.go
# change ***** to your info

make db && make redis && make migration

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
# change ***** to your info
./deploy.sh
```


