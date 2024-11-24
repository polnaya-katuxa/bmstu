plate_group(wf).
plate_group(pm).
plate_group(twothree).
plate_group(ngp).
plate_group(cnmx-sm).
plate_group(knux).
plate_group(vcex).
plate_group(rngn).


% группа пластинок, список доступных для нее групп металлов, износостойкойсть для группы металлов
plate_properties(wf, ["M", "P", "K"], [20, 30, 40]).
plate_properties(pm, ["P"], [30]).
plate_properties(twothree, ["M", "S", "N"], [30, 20, 40]).
plate_properties(ngp, ["S"], [15]).
plate_properties(cnmx-sm, ["S", "M"], [10, 50]).
plate_properties(knux, ["P", "M", "K", "S"], [20, 20, 30, 40]).
plate_properties(vcex, ["P", "M", "K", "S", "N"], [30, 20, 30, 40, 50]).
plate_properties(rngn, ["H", "K"], [20, 30]).

plate(wf, "CNMG090304-WF", ap(50, 30, 150), fn(15, 5, 25), cost(625)).
plate(wf, "CNMG090308-WF", ap(100, 30, 200), fn(30, 10, 50), cost(460)).
plate(wf, "CNMG120404-WF", ap(40, 30, 300), fn(15, 5, 25), cost(755)).
plate(wf, "CNMG120408-WF", ap(100, 30, 400), fn(30, 10, 50), cost(345)).
plate(wf, "CNMG120412-WF", ap(150, 40, 400), fn(50, 20, 60), cost(834)).
plate(wf, "DNMX110404-WF", ap(100, 20, 150), fn(20, 8, 30), cost(659)).
plate(wf, "DNMX110408-WF", ap(100, 20, 300), fn(30, 10, 40), cost(358)).
plate(wf, "DNMX150404-WF", ap(80, 20, 300), fn(20, 8, 30), cost(319)).
plate(wf, "DNMX150408-WF", ap(150, 20, 300), fn(30, 10, 40), cost(979)).
plate(wf, "DNMX150412-WF", ap(150, 40, 350), fn(40, 15, 55), cost(300)).
plate(wf, "DNMX150604-WF", ap(80, 20, 300), fn(20, 8, 30), cost(908)).
plate(wf, "DNMX150608-WF", ap(150, 20, 300), fn(30, 10, 40), cost(893)).
plate(wf, "DNMX150612-WF", ap(150, 40, 350), fn(40, 15, 55), cost(446)).
plate(wf, "TNMX160404-WF", ap(100, 20, 300), fn(20, 8, 30), cost(307)).
plate(wf, "TNMX160408-WF", ap(150, 20, 300), fn(30, 10, 40), cost(956)).
plate(wf, "WNMG060404-WF", ap(40, 30, 200), fn(15, 5, 25),  cost(681)).
plate(wf, "WNMG060408-WF", ap(100, 30, 300), fn(30, 10, 50), cost(177)).
plate(wf, "WNMG080404-WF", ap(40, 30, 300), fn(15, 5, 25), cost(737)).
plate(wf, "WNMG080408-WF", ap(100, 30, 400), fn(30, 10, 50), cost(353)).
plate(wf, "WNMG080412-WF", ap(150, 40, 400), fn(50, 20, 60), cost(739)).

plate(pm, "CNMG090304-PM", ap(200, 40, 400), fn(20, 10, 30), cost(809)).
plate(pm, "CNMG090308-PM", ap(200, 50, 400), fn(30, 15, 50), cost(513)).
plate(pm, "CNMG120404-PM", ap(300, 40, 550), fn(20, 10, 30), cost(335)).
plate(pm, "CNMG120408-PM", ap(300, 50, 550), fn(30, 15, 50), cost(205)).
plate(pm, "CNMG120412-PM", ap(300, 80, 550), fn(35, 18, 60), cost(170)).
plate(pm, "CNMG120416-PM", ap(300, 100, 550), fn(40, 23, 65), cost(930)).
plate(pm, "CNMG160608-PM", ap(400, 50, 720), fn(30, 15, 50), cost(884)).
plate(pm, "CNMG160612-PM", ap(400, 80, 720), fn(35, 18, 60), cost(630)).
plate(pm, "CNMG160616-PM", ap(400, 100, 720), fn(40, 23, 65), cost(145)).
plate(pm, "CNMG190608-PM", ap(400, 50, 860), fn(30, 15, 50), cost(914)).
plate(pm, "CNMG190612-PM", ap(400, 80, 860), fn(35, 18, 60), cost(659)).
plate(pm, "CNMG190616-PM", ap(400, 100, 860), fn(40, 23, 65), cost(934)).
plate(pm, "DNMG110404-PM", ap(200, 40, 500), fn(20, 10, 30), cost(604)).
plate(pm, "DNMG110408-PM", ap(200, 50, 500), fn(30, 15, 50), cost(143)).
plate(pm, "DNMG110412-PM", ap(200, 80, 500), fn(35, 18, 50), cost(560)).
plate(pm, "DNMG150404-PM", ap(300, 40, 600), fn(20, 10, 30), cost(837)).
plate(pm, "DNMG150408-PM", ap(300, 50, 600), fn(30, 15, 50), cost(849)).
plate(pm, "DNMG150412-PM", ap(300, 80, 600), fn(35, 18, 60), cost(343)).
plate(pm, "DNMG150604-PM", ap(300, 40, 600), fn(20, 10, 30), cost(944)).
plate(pm, "DNMG150608-PM", ap(300, 50, 600), fn(30, 15, 50), cost(824)).
plate(pm, "DNMG150612-PM", ap(300, 80, 600), fn(35, 18, 60), cost(930)).
plate(pm, "DNMG150616-PM", ap(300, 100, 600), fn(40, 23, 65), cost(681)).
plate(pm, "SNMG090304-PM", ap(200, 40, 450), fn(20, 10, 30), cost(100)).
plate(pm, "SNMG090308-PM", ap(200, 50, 450), fn(30, 15, 50), cost(890)).
plate(pm, "SNMG120404-PM", ap(300, 40, 600), fn(20, 10, 30), cost(870)).
plate(pm, "SNMG120408-PM", ap(300, 50, 600), fn(30, 15, 50), cost(572)).
plate(pm, "SNMG120412-PM", ap(300, 80, 600), fn(35, 18, 60), cost(416)).
plate(pm, "SNMG120416-PM", ap(300, 100, 600), fn(40, 23, 65), cost(885)).
plate(pm, "SNMG150612-PM", ap(400, 80, 750), fn(35, 18, 60), cost(867)).
plate(pm, "SNMG150616-PM", ap(400, 100, 750), fn(40, 23, 65), cost(454)).
plate(pm, "TNMG160404-PM", ap(300, 40, 500), fn(20, 10, 30), cost(159)).
plate(pm, "TNMG160408-PM", ap(300, 50, 500), fn(30, 15, 50), cost(299)).
plate(pm, "TNMG160412-PM", ap(300, 80, 500), fn(35, 18, 60), cost(673)).
plate(pm, "TNMG220404-PM", ap(400, 40, 660), fn(20, 10, 30), cost(888)).
plate(pm, "TNMG220408-PM", ap(400, 50, 660), fn(30, 15, 50), cost(678)).
plate(pm, "TNMG220412-PM", ap(400, 80, 660), fn(35, 18, 60), cost(798)).
plate(pm, "TNMG220416-PM", ap(400, 100, 660), fn(40, 23, 65), cost(788)).
plate(pm, "VNMG160408-PM", ap(200, 50, 400), fn(30, 15, 50), cost(209)).
plate(pm, "VNMG160412-PM", ap(200, 80, 400), fn(35, 18, 60), cost(689)).
plate(pm, "WNMG060408-PM", ap(200, 50, 300), fn(30, 15, 50), cost(737)).
plate(pm, "WNMG060412-PM", ap(200, 80, 300), fn(35, 18, 60), cost(481)).
plate(pm, "WNMG080408-PM", ap(250, 50, 400), fn(30, 15, 50), cost(913)).
plate(pm, "WNMG080412-PM", ap(250, 80, 400), fn(35, 18, 60), cost(530)).
plate(pm, "WNMG080416-PM", ap(300, 100, 400), fn(40, 23, 65), cost(497)).

plate(twothree, "CNMG120404-23", ap(150, 20, 360), fn(14, 10, 18), cost(842)).
plate(twothree, "CNMG120408-23", ap(240, 50, 360), fn(18, 13, 24), cost(969)).
plate(twothree, "CNMG120412-23", ap(240, 40, 360), fn(22, 16, 29), cost(848)).
plate(twothree, "CNMG160608-23", ap(290, 40, 430), fn(21, 16, 29), cost(544)).
plate(twothree, "CNMG160612-23", ap(290, 50, 430), fn(26, 19, 34), cost(625)).
plate(twothree, "CNMG160616-23", ap(290, 60, 430), fn(32, 24, 43), cost(966)).
plate(twothree, "CNMG190608-23", ap(300, 150, 800), fn(35, 25, 50), cost(369)).
plate(twothree, "CNMG190612-23", ap(320, 60, 480), fn(29, 21, 38), cost(189)).
plate(twothree, "CNMG190616-23", ap(300, 150, 800), fn(40, 35, 65), cost(684)).
plate(twothree, "DNMG110404-23", ap(200, 40, 400), fn(15, 10, 20), cost(678)).
plate(twothree, "DNMG150404-23", ap(150, 20, 360), fn(14, 10, 18), cost(444)).
plate(twothree, "DNMG150408-23", ap(240, 40, 360), fn(18, 13, 24), cost(408)).
plate(twothree, "DNMG150412-23", ap(240, 40, 360), fn(22, 16, 29), cost(538)).
plate(twothree, "DNMG150604-23", ap(150, 20, 360), fn(14, 10, 18), cost(621)).
plate(twothree, "DNMG150608-23", ap(240, 40, 360), fn(18, 13, 50), cost(128)).
plate(twothree, "DNMG150612-23", ap(240, 40, 360), fn(22, 16, 29), cost(172)).
plate(twothree, "SNMG120404-23", ap(150, 20, 360), fn(14, 10, 18), cost(350)).
plate(twothree, "SNMG120408-23", ap(240, 40, 360), fn(18, 13, 24), cost(339)).
plate(twothree, "SNMG120412-23", ap(240, 40, 360), fn(22, 16, 29), cost(352)).
plate(twothree, "SNMG150612-23", ap(300, 150, 750), fn(35, 30, 60), cost(933)).
plate(twothree, "SNMG150616-23", ap(300, 150, 750), fn(40, 35, 65), cost(442)).
plate(twothree, "SNMG190612-23", ap(320, 60, 480), fn(29, 21, 38), cost(843)).
plate(twothree, "SNMG190616-23", ap(300, 150, 800), fn(40, 35, 65), cost(493)).
plate(twothree, "TNMG160308-23", ap(480, 150, 800), fn(50, 40, 56), cost(435)).
plate(twothree, "TNMG160404-23", ap(200, 20, 300), fn(11, 8, 15), cost(263)).
plate(twothree, "TNMG160408-23", ap(200, 30, 300), fn(15, 11, 20), cost(368)).
plate(twothree, "TNMG160412-23", ap(200, 40, 300), fn(18, 13, 24), cost(278)).
plate(twothree, "TNMG220408-23", ap(240, 40, 360), fn(18, 13, 24), cost(372)).
plate(twothree, "TNMG220412-23", ap(240, 40, 360), fn(22, 16, 29), cost(716)).
plate(twothree, "VNMG160404-23", ap(200, 40, 400), fn(15, 10, 20), cost(430)).
plate(twothree, "VNMG160408-23", ap(250, 50, 400), fn(20, 15, 25), cost(843)).
plate(twothree, "WNMG060404-23", ap(200, 50, 300), fn(15, 10, 30), cost(269)).
plate(twothree, "WNMG060408-23", ap(250, 70, 400), fn(25, 20, 35), cost(315)).
plate(twothree, "WNMG080404-23", ap(250, 50, 400), fn(15, 10, 30), cost(176)).
plate(twothree, "WNMG080408-23", ap(250, 70, 400), fn(25, 20, 35), cost(448)).
plate(twothree, "WNMG080412-23", ap(250, 100, 400), fn(30, 25, 35), cost(1000)).

plate(ngp, "CNGP120404", ap(60, 10, 130), fn(11, 6, 15), cost(521)).
plate(ngp, "CNGP120408", ap(60, 20, 130), fn(17, 10, 25), cost(423)).
plate(ngp, "DNGP150604", ap(20, 10, 30), fn(10, 5, 15), cost(552)).
plate(ngp, "DNGP150608", ap(30, 20, 50), fn(17, 10, 25), cost(179)).
plate(ngp, "VNGP160404-DS", ap(20, 10, 30), fn(10, 5, 15), cost(440)).
plate(ngp, "VNGP160408-DS", ap(30, 20, 50), fn(17, 10, 25), cost(438)).

plate(cnmx-sm, "CNMX1204A1-SM", ap(100, 50, 150), fn(25, 13, 35), cost(966)).
plate(cnmx-sm, "CNMX1204A2-SM", ap(200, 50, 250), fn(25, 13, 35), cost(317)).

plate(knux, "KNUX160405FL12", ap(400, 150, 600), fn(30, 25, 35), cost(505)).
plate(knux, "KNUX160405FR12", ap(400, 150, 600), fn(30, 25, 35), cost(415)).
plate(knux, "KNUX160405L11", ap(300, 100, 600), fn(30, 20, 35), cost(635)).
plate(knux, "KNUX160405L12", ap(400, 150, 600), fn(30, 25, 35), cost(423)).
plate(knux, "KNUX160405R11", ap(300, 100, 600), fn(30, 20, 35), cost(741)).
plate(knux, "KNUX160405R12", ap(400, 150, 600), fn(30, 25, 35), cost(403)).
plate(knux, "KNUX160410FL12", ap(400, 150, 600), fn(50, 40, 55), cost(429)).
plate(knux, "KNUX160410FR12", ap(400, 150, 600), fn(50, 40, 55), cost(453)).
plate(knux, "KNUX160410L11", ap(300, 150, 600), fn(40, 30, 60), cost(289)).
plate(knux, "KNUX160410L12", ap(400, 150, 600), fn(50, 40, 70), cost(547)).
plate(knux, "KNUX160410R11", ap(300, 150, 600), fn(40, 30, 60), cost(979)).
plate(knux, "KNUX160410R12", ap(400, 150, 600), fn(50, 40, 70), cost(442)).
plate(knux, "KNUX160415FL13", ap(400, 100, 800), fn(60, 50, 70), cost(392)).
plate(knux, "KNUX160415FR13", ap(400, 100, 800), fn(60, 50, 70), cost(848)).

plate(vcex, "VCEX110300L-F", ap(100, 3, 400), fn(5, 5, 20), cost(174)).
plate(vcex, "VCEX110300R-F", ap(100, 3, 400), fn(5, 5, 20), cost(604)).
plate(vcex, "VCEX110301L-F", ap(100, 5, 400), fn(10, 5, 30), cost(560)).
plate(vcex, "VCEX110301R-F", ap(100, 5, 400), fn(10, 5, 30), cost(270)).

plate(rngn, "RNGN090300T01020", ap(90, 10, 270), fn(65, 13, 244), cost(711)).
plate(rngn, "RNGN090300E", ap(90, 10, 270), fn(65, 13, 244), cost(682)).
plate(rngn, "RNGN120400T01020", ap(120, 10, 360), fn(81, 13, 282), cost(683)).
plate(rngn, "RNGN120400T02520", ap(360, 10, 480), fn(56, 56, 24, 676), cost(345)).
plate(rngn, "RNGN120700T01020", ap(120, 10, 360), fn(81, 13, 282), cost(144)).
plate(rngn, "RNGN120700E", ap(120, 10, 360), fn(81, 19, 451), cost(459)).
plate(rngn, "RNGN120700K15015", ap(120, 10, 360), fn(81, 13, 282), cost(997)).
plate(rngn, "RNGN120700T02520", ap(360, 10, 480), fn(56, 24, 676), cost(217)).
plate(rngn, "RNGN120700T15015", ap(120, 10, 360), fn(81, 13, 282), cost(777)).
plate(rngn, "RNGN150700T01020", ap(150, 10, 450), fn(81, 19, 504), cost(775)).
plate(rngn, "RNGN150700T02520", ap(150, 10, 450), fn(81, 13, 315), cost(639)).
plate(rngn, "RNGN150700T20015", ap(150, 10, 450), fn(81, 13, 315), cost(667)).
plate(rngn, "RNGN190700T01020", ap(190, 10, 570), fn(95, 18, 552), cost(671)).
plate(rngn, "RNGN190700K20015", ap(190, 10, 570), fn(95, 13, 345), cost(286)).
plate(rngn, "RNGN190700T20015", ap(190, 10, 570), fn(95, 13, 345), cost(210)).
plate(rngn, "RNGN250700K20015", ap(250, 10, 750), fn(96, 13, 398), cost(653)).
plate(rngn, "RNGN250700T20015", ap(250, 10, 750), fn(96, 13, 398), cost(991)).
