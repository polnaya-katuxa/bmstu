module.exports = {
  async up(db, client) {
      await db.createCollection( "users",
        {
          validator: { $or:
                [
                  { mail: { $regex: /@*\.*$/ } },
                  { balance: { $gte: 0 } }
                ]
          }
        }
    );
       db.collection('users').createIndex({ login: 1 }, { unique: true, partialFilterExpression: { login: { $exists: true} } });
       db.collection('users').createIndex({ password: 1 },  { partialFilterExpression: { password: { $exists: true } } });
       db.collection('users').createIndex({ mail: 1 },{ partialFilterExpression: { mail: { $exists: true } } });
       db.collection('users').createIndex({ enter_time: 1 }, { partialFilterExpression: { enter_time: { $exists: true } } });

     db.createCollection( "reaction_types");
       db.collection('reaction_types').createIndex({ icon: 1 },  { partialFilterExpression: { icon: { $exists: true } } });

     db.createCollection( "limits",
        {
          validator: { $or:
                [
                  { value: { $gte: 0 } },
                  { bonus: { $gte: 0 } }
                ]
          }
        }
    );
       db.collection('limits').createIndex({ value: 1 }, { unique: true, partialFilterExpression: { value: { $exists: true } } });
       db.collection('limits').createIndex({ bonus: 1 },  { partialFilterExpression: { bonus: { $exists: true } } });

     db.createCollection( "posts",
        {
          validator: { $or:
                [
                  { perms: { $in: [ "free", "paid" ] } }
                ]
          }
        }
    );
       db.collection('posts').createIndex({ content: 1 }, { partialFilterExpression: { content: { $exists: true } } });
       db.collection('posts').createIndex({ public_time: 1 },  { partialFilterExpression: { public_time: { $exists: true } } });

     db.createCollection( "reactions");

     db.createCollection( "subscriptions");

     db.createCollection( "balance_transactions");
       db.collection('balance_transactions').createIndex({ reason: 1 }, { partialFilterExpression: { reason: { $exists: true } } });
       db.collection('balance_transactions').createIndex({ time: 1 }, { partialFilterExpression: { time: { $exists: true } } });
       db.collection('balance_transactions').createIndex({ amount: 1 }, { partialFilterExpression: { amount: { $exists: true } } });

     db.createCollection( "comments");
       db.collection('comments').createIndex({ public_time: 1 }, { partialFilterExpression: { public_time: { $exists: true } } });
       db.collection('comments').createIndex({ content: 1 }, { partialFilterExpression: { content: { $exists: true } } });
  },

  async down(db, client) {
       db.collection('comments').dropIndexes();
       db.collection('comments').drop();

       db.collection('balance_transactions').dropIndexes();
       db.collection('balance_transactions').drop();

       db.collection('subscriptions').dropIndexes();
       db.collection('subscriptions').drop();

       db.collection('reactions').dropIndexes();
       db.collection('reactions').drop();

       db.collection('posts').dropIndexes();
       db.collection('posts').drop();

       db.collection('limits').dropIndexes();
       db.collection('limits').drop();

       db.collection('reaction_types').dropIndexes();
       db.collection('reaction_types').drop();

       db.collection('users').dropIndexes();
       db.ollection('users').drop();
  }
};
