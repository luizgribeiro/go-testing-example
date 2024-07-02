count=0

while [ $count -le 30 ]; do
    liquibase --changelog-file ./changelog/db.changelog.yaml --url jdbc:postgresql://postgres:5432/mydatabase --username myuser --password mypassword update

    if [ "$?" -eq 0 ]; then
        echo "Migration executed!"
        break
    fi

    echo "Migration failed --- sleeping"

    sleep 2
done;
