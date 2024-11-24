module.exports = {
  async up(db, client) {
      const { Binary, ObjectID } = require('mongodb');

    await db.collection('limits').insertMany(
        [
          { _id: new Binary(Buffer.from('GbmzP91jRzOCXntgahIR9A==', 'base64')), value: 5, bonus: 10}, //BinData(0, "GbmzP91jRzOCXntgahIR9A=="), value: 5, bonus: 10},
          { _id: new Binary(Buffer.from('9dt5j1PvRPWMaIgnl/rBEg==', 'base64')), value: 10, bonus: 20}, //BinData(0, "9dt5j1PvRPWMaIgnl/rBEg=="), value: 10, bonus: 20},
          { _id: new Binary(Buffer.from('1iLfpauaRwKdKWzXxiWb0w==', 'base64')), value: 20, bonus: 30}, //BinData(0, "1iLfpauaRwKdKWzXxiWb0w=="), value: 20, bonus: 30},
          { _id: new Binary(Buffer.from('N1vl1DSnQsGy36/rheRlNA==', 'base64')), value: 50, bonus: 40}, //BinData(0, "N1vl1DSnQsGy36/rheRlNA=="), value: 50, bonus: 40},
          { _id: new Binary(Buffer.from('BkskWvMMSpGL4DpCAug9wA==', 'base64')), value: 100, bonus: 50}, //BinData(0, "BkskWvMMSpGL4DpCAug9wA=="), value: 100, bonus: 50},
          { _id: new Binary(Buffer.from('sclxmOSAQbGu8d03dND20Q==', 'base64')), value: 2147483647, bonus: 0}, //BinData(0, "sclxmOSAQbGu8d03dND20Q=="), value: 2147483647, bonus: 0}
        ]
    );

    await db.collection('reaction_types').insertMany(
        [
          { _id: new Binary(Buffer.from('mYNbkgYBQYScVWqc1az4Fg==', 'base64')), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/red-heart.png"}, //BinData(0, "mYNbkgYBQYScVWqc1az4Fg=="), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/red-heart.png"},
          { _id: new Binary(Buffer.from('EO0+0aYaQDa6WzG9Ozry9Q==', 'base64')), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/fire.png"}, //BinData(0, "EO0+0aYaQDa6WzG9Ozry9Q=="), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/fire.png"},
          { _id: new Binary(Buffer.from('RySLEt7iRxWiEgg9fzAjmQ==', 'base64')), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/disguised-face.png"}, //BinData(0, "RySLEt7iRxWiEgg9fzAjmQ=="), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/disguised-face.png"},
          { _id: new Binary(Buffer.from('U9SS2d5nQmmcQ7FrxOJ3CQ==', 'base64')), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/broken-heart.png"}, //BinData(0, "U9SS2d5nQmmcQ7FrxOJ3CQ=="), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/broken-heart.png"},
          { _id: new Binary(Buffer.from('4syPL3/rQtOqaekjhYAdUQ==', 'base64')), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/poop.png"}, //BinData(0, "4syPL3/rQtOqaekjhYAdUQ=="), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/poop.png"},
          { _id: new Binary(Buffer.from('JnTnpvIkRqu7afFxLO7X4A==', 'base64')), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/sneezing-face.png"} //BinData(0, "JnTnpvIkRqu7afFxLO7X4A=="), icon: "https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/sneezing-face.png"}
        ]
    );
  },

  async down(db, client) {
      await db.collection('reaction_types').remove();
      await db.collection('limits').remove();
  }
};
