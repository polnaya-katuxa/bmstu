#ifndef DRAWING_H
#define DRAWING_H

#include "figure.h"
#include "lightsource.h"
#include "rotation.h"
#include "sight.h"
#include <QImage>
#include <QPainter>
#include <QVector3D>

class Drawing {
public:
    Drawing(int depth = 3);
    Drawing(std::vector<std::shared_ptr<Figure>> figures, int depth = 3);
    Drawing(std::vector<std::shared_ptr<Figure>> figures, int canvasHeight, int canvasWidth, int depth = 3);
    Drawing(int canvasHeight, int canvasWidth, int depth = 3);

    bool loadFile(QString filename);
    bool loadFigures(std::string filename);
    bool loadSphere(QVector3D cen, double rad, QColor dC, QColor sC, QVector4D alb, double rCoef, double sCoef);

    std::tuple<bool, double, QVector3D, int> sceneIntersect(const sight_t& sight);
    QColor castRay(const sight_t& sight, const int& depth);
    std::shared_ptr<QImage> drawFigures(int numThreads);
    void rotate(QVector3D angles);
    void initDefaultLightSources();

    void drawThread(QVector3D &preLim, QVector3D &lim, sight_t sight, std::vector<std::vector<uint>> &image);

private:
    std::vector<std::shared_ptr<Figure>> figures;
    std::vector<LightSource> lightSources;

    int canvasHeight = 900;
    int canvasWidth = 1600;
    int depth = 3;
};

#endif // DRAWING_H
