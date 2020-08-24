#!/bin/bash

echo -e "[INFO] inicializando el setup de su proyecto\n"

PROJECT_LANG=""
PS3="Q: Cuál es su lenguaje de preferencia? (default: python - recomendado) "
#ls src |xargs echo "quit"
options=("go" "python" "nodejs" "quit")
PROJECT_LANG=python
select opt in "${options[@]}"
do
    case $opt in
        "go")
            echo "Eligió $REPLY: $opt"
            PROJECT_LANG=go
            break
            ;;
        "python")
            echo "Eligió $REPLY: $opt"
            PROJECT_LANG=python
            break
            ;;
        "nodejs")
            echo "Eligió $REPLY: $opt"
            PROJECT_LANG=nodejs
            break
            ;;

        "Quit")
            break
            ;;
        *) echo "opción inválida $REPLY";;
    esac
done

echo -e "\nQ: Tiene un repo en docker hub? de ser asi cual es? (usualmente este es su Docker ID)"
read docker_repo

echo -e "\nQ: Que dominio (DNS local) le gustaria usar? [default: example.com]"
read domain

echo "# [autogenerado] esto puede ser modificado en todo momento" > .env
echo -e "\n[INFO] Polularemos su .env con los siguientes datos"
echo """
PROJECT_LANG=${PROJECT_LANG}
DOCKER_REPO=${docker_repo}
MY_DOMAIN=${domain:-'example.com'}
""" | tee -a .env
