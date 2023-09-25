-- +goose Up
-- +goose StatementBegin
create table cat (
                     breed text,
                     fluffiness int,
                     image text
);

insert into cat (breed, fluffiness, image) values
    ('Канадский сфинкс', 0, 'https://mozhaiskiy-gazeta.ru/files/data/user/elena/files/2021.10.05-1633446567.0222_nyjh2rj00ju.jpg'),
    ('Сиамская', 1005,'https://img.freepik.com/premium-photo/siamese-cat-in-front-of-white-background_87557-23043.jpg'),
    ('Сибирская', 18005, 'https://avatars.dzeninfra.ru/get-zen_doc/1930013/pub_61e977c748d6ed57787a77de_61e977d8b515ee29e433b68c/scale_1200'),
    ('Мейн-кун', 12900,'https://kotopediya.su/wp-content/uploads/2017/06/mein-kun3.jpg'),
    ('Персидская', 20000, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT2wbPZHVyq6x8KFQoxJWUaGCtlbknvc1lyvQ&usqp=CAU'),
    ('Ориентальная', 424, 'https://i.ytimg.com/vi/01tGD2YuKsQ/maxresdefault.jpg'),
    ('Курильский бобтейл', 1580, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR6Y83EkdFHSMYW4X4UFdsjqWen6h3MifKPeA&usqp=CAU'),
    ('Британская', 2043, 'https://www.meme-arsenal.com/memes/fc06bda460d74194f13d633368e804f3.jpg'),
    ('Рэгдолл', 16500, 'https://cs12.pikabu.ru/post_img/2022/07/02/11/1656787574162939768.jpg'),
    ('Донской cфинкс', 1, 'https://www.jvlife.ru/system/images/contents/000/008/644/medium_cropped/AdobeStock_74523649.jpg?1553264540'),
    ('Петерболд', 5, 'https://www.cats-british.ru/files/breeds/kotenok_peterbold.jpeg'),
    ('Украинский левкой', 107, 'https://skischool-nso.ru/wp-content/uploads/2013/01/Ukrainian-Levkoy.jpg'),
    ('Йоркская шоколадная', 13765, 'https://kotology.ru/wp-content/uploads/2016/11/2659856071.jpg'),
    ('Рагамаффин', 6897, 'https://prohvost.club/wp-content/uploads/2018/03/ragamaffin-na-rukah-u-hozyayki-posle-pobedy-na-vystavke.jpg'),
    ('Американский кёрл', 11922, 'https://static.wixstatic.com/media/973af9_d8d5d6ea3df6479daf84bd431f0f51e3~mv2.jpg/v1/fill/w_640,h_580,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/973af9_d8d5d6ea3df6479daf84bd431f0f51e3~mv2.jpg'),
    ('Нибелунг', 8777, 'https://irecommend.ru/sites/default/files/imagecache/copyright1/user-images/546937/QFfrkMTiBVMqiTxTvweUA.JPG'),
    ('Турецкая ангора', 13550,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQMX-L7X4rP22OO4qSlFWgU_BJZzADShGpRBQ&usqp=CAU'),
    ('Бирманская', 10021, 'https://catpeople.ru/bried/birmanskaya/valda06112014-103.jpg'),
    ('Абиссинская', 865, 'https://irecommend.ru/sites/default/files/imagecache/copyright1/user-images/2062240/r0k3tEiKsdoVIyqepTa7Tw.jpg'),
    ('Бенгальская', 5733, 'https://bonifacyholidays.ru/assets/images/bengal2.jpg'),
    ('Бомбейская', 999, 'https://brothers-smaller.ru/wp-content/uploads/2015/03/032115_1334_3.jpg'),
    ('Саванна', 653, 'https://avatars.mds.yandex.net/i?id=98c7e9ff05129d7e4dbbc0a966c919c9-5268158-images-thumbs&n=13'),
    ('Пиксибоб', 3211, 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/48/Alsoomse_Pixiebob_Whisperer.jpg/640px-Alsoomse_Pixiebob_Whisperer.jpg'),
    ('Манчкин', 9782, 'https://img-fotki.yandex.ru/get/15502/147960462.0/0_11f1f3_8586e334_L.jpg'),
    ('Корниш-рекс', 112, 'https://www.purina.ru/sites/default/files/2021-09/Корниш-рекс%205-min.jpg');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cat cascade;
-- +goose StatementEnd
