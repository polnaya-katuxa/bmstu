#include "drawing.h"
#include "mainwindow.h"
#include "polyhedron.h"
#include "sphere.h"

#include <QApplication>
#include <QTest>
#include <getopt.h>
#include <iostream>
#include <unistd.h>

#include "test_rotation.h"

#define NUM_THREADS 8

int main(int argc, char* argv[])
{
    if (argc == 1) {
        QApplication a(argc, argv);
        MainWindow w;
        w.show();
        return a.exec();
    }

    int depth = -1;
    int angle = 0;
    int threads = 0;
    bool measure = false;
    char* iFile = nullptr;
    char* oFile = nullptr;

    int res = 0;
    while ((res = getopt(argc, argv, "d:i:o:r:tmn:")) != -1) {
        switch (res) {
        case 'd':
            depth = std::stoi(optarg);
            break;
        case 'i':
            iFile = optarg;
            break;
        case 'o':
            oFile = optarg;
            break;
        case 'r':
            angle = std::stoi(optarg);
            break;
        case 'n':
            threads = std::stoi(optarg);
            break;
        case 't':
            return QTest::qExec(new TestRotation);
        case 'm':
            measure = true;
            break;
        case '?':
            return EXIT_FAILURE;
        }
    }
    if (depth <= 0 || !iFile)
        return EXIT_FAILURE;

    if (measure) {
        using std::chrono::duration_cast;
        using std::chrono::microseconds;

        auto end = std::chrono::steady_clock::now();
        auto start = std::chrono::steady_clock::now();

        start = std::chrono::steady_clock::now();

        Drawing picture = Drawing(depth);
        picture.initDefaultLightSources();

        if (!picture.loadFile(QString(iFile)))
            return EXIT_FAILURE;

        picture.drawFigures(threads);

        end = std::chrono::steady_clock::now();

        std::cout << (double)duration_cast<microseconds>(end - start).count() / 1000 << "\n";

        return EXIT_SUCCESS;
    }

    if (!oFile)
        return EXIT_FAILURE;

    Drawing picture = Drawing(depth);
    picture.initDefaultLightSources();

    if (!picture.loadFile(QString(iFile)))
        return EXIT_FAILURE;

    picture.rotate(QVector3D(0, angle, 0));

    std::shared_ptr<QImage> curImage = picture.drawFigures(threads);
    curImage->save(oFile);

    return 0;
}
