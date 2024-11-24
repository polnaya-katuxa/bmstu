db.createUser(
    {
        user: "postby",
        pwd: "password",
        roles: [
            {
                role: "readWrite",
                db: "postby"
            }
        ]
    }
);
