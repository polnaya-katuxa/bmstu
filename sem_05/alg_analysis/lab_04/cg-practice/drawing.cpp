#include <QMessageBox>
#include <QVector3D>
#include <thread>

#include "drawing.h"
#include "objloader.h"
#include "polyhedron.h"
#include "sphere.h"

Drawing::Drawing(int depth)
    : depth(depth)
{
}

Drawing::Drawing(std::vector<std::shared_ptr<Figure>> figures, int depth)
    : figures(figures)
    , depth(depth)
{
}

Drawing::Drawing(std::vector<std::shared_ptr<Figure>> figures, int canvasHeight, int canvasWidth, int depth)
    : figures(figures)
    , canvasHeight(canvasHeight)
    , canvasWidth(canvasWidth)
    , depth(depth)
{
}

Drawing::Drawing(int canvasHeight, int canvasWidth, int depth)
    : canvasHeight(canvasHeight)
    , canvasWidth(canvasWidth)
    , depth(depth)
{
}

void readQVector3D(QVector3D& vec, std::ifstream& file)
{
    double coord;

    file >> coord;
    vec.setX(coord);
    file >> coord;
    vec.setY(coord);
    file >> coord;
    vec.setZ(coord);
}

void readQColor(QColor& col, std::ifstream& file)
{
    int elem;

    file >> elem;
    col.setRed(elem);
    file >> elem;
    col.setGreen(elem);
    file >> elem;
    col.setBlue(elem);
}

void readQVector4D(QVector4D& vec, std::ifstream& file)
{
    double coord;

    file >> coord;
    vec.setX(coord);
    file >> coord;
    vec.setY(coord);
    file >> coord;
    vec.setZ(coord);
    file >> coord;
    vec.setW(coord);
}

bool Drawing::loadFile(QString filename)
{
    std::ifstream file;
    std::string line;
    int num = 0;

    file.open(filename.toStdString());

    if (file.is_open()) {
        while (file >> line) {
            num++;
            if (line == "O") {
                file >> line;
                if (!loadFigures(line))
                    return false;
            } else if (line == "A") {
                QVector3D cen;
                double rad;
                QColor dC, sC;
                QVector4D alb;
                double rCoef, sCoef;

                readQVector3D(cen, file);
                file >> rad;
                readQColor(dC, file);
                readQColor(sC, file);
                readQVector4D(alb, file);
                file >> rCoef;
                file >> sCoef;

                if (!loadSphere(cen, rad, dC, sC, alb, rCoef, sCoef))
                    return false;
            } else
                return false;
        }
        if (!num)
            return false;
        file.close();
    } else
        return false;

    return true;
}

bool Drawing::loadFigures(std::string filename)
{
    objl::Loader loader;

    bool loaded = loader.LoadFile(filename);

    if (loaded) {
        size_t minVertPos = 0;
        size_t minPolPos = 0;

        for (size_t i = 0; i < loader.LoadedMeshes.size(); i++) {
            objl::Mesh curMesh = loader.LoadedMeshes[i];

            Polyhedron p = Polyhedron();
            std::vector<QVector3D> points;
            std::vector<Polygon> polygons;

            for (size_t j = minVertPos; j < curMesh.Positions.size(); j++) {
                QVector3D point = QVector3D(curMesh.Positions[j].X, curMesh.Positions[j].Y, curMesh.Positions[j].Z);
                points.push_back(point);
            }

            minVertPos = curMesh.Positions.size();

            for (size_t j = minPolPos; j < curMesh.Indexes.size(); j += 3) {
                Polygon pol = Polygon({ curMesh.Indexes[j] - 1, curMesh.Indexes[j + 1] - 1, curMesh.Indexes[j + 2] - 1 }, points);
                polygons.push_back(pol);
            }

            minPolPos = curMesh.Indexes.size();

            p.setPoints(points);
            p.setPolygons(polygons);

            p.setMaterial(QVector4D(curMesh.MeshMaterial.Ka.X, curMesh.MeshMaterial.Ka.Y, curMesh.MeshMaterial.Ka.Z, curMesh.MeshMaterial.d),
                QColor(curMesh.MeshMaterial.Kd.X * 255, curMesh.MeshMaterial.Kd.Y * 255, curMesh.MeshMaterial.Kd.Z * 255),
                QColor(curMesh.MeshMaterial.Ks.X * 255, curMesh.MeshMaterial.Ks.Y * 255, curMesh.MeshMaterial.Ks.Z * 255),
                curMesh.MeshMaterial.Ni, curMesh.MeshMaterial.Ns);

            figures.push_back(std::make_shared<Polyhedron>(p));
        }
        return true;
    }
    return false;
}

bool Drawing::loadSphere(QVector3D cen, double rad, QColor dC, QColor sC, QVector4D alb, double rCoef, double sCoef)
{
    Sphere s = Sphere(cen, rad);
    s.setMaterial(alb, dC, sC, rCoef, sCoef);

    figures.push_back(std::make_shared<Sphere>(s));

    return true;
}

void Drawing::initDefaultLightSources()
{
    lightSources = {
        LightSource(QVector3D(600, 250, 280), 0.9),
        LightSource(QVector3D(600, 1000, 1000), 0.8),
        LightSource(QVector3D(-800, -800, 314), 0.7),
        LightSource(QVector3D(900, 300, 0), 1),
    };
}

QVector3D reflect(const QVector3D& I, const QVector3D& N)
{
    return I - N * 2.0 * QVector3D::dotProduct(I, N);
}

QVector3D refract(const QVector3D& I, const QVector3D& N, const double& t, const double i = 1.f)
{
    double cosI = -std::max<double>(-1., std::min<double>(1., QVector3D::dotProduct(I, N)));
    if (cosI < 0)
        return refract(I, N * (-1), i, t);
    double c = i / t;
    double k = 1 - c * c * (1 - cosI * cosI);
    return k < 0 ? QVector3D(1, 0, 0) : I * c + N * (c * cosI - std::sqrt(k));
}

double getColorElem(const double& in, const double& specIn, const QVector4D& alb, const int& figCol, const int& ligCol, const int& refCol, const int refractCol)
{
    double c = in * figCol * alb.x() + specIn * ligCol * alb.y() + refCol * alb.z() + refractCol * alb.w();
    if (c > 255)
        c = 255;

    return c;
}

QColor getColor(const double& intensity, const double& specIntensity, const QVector4D& alb, const QColor& cF, const QColor& cL, const QColor& cR, const QColor& cRef)
{
    QColor color;
    int rF, gF, bF, rL, gL, bL, rR, gR, bR, rRef, gRef, bRef;
    cF.getRgb(&rF, &gF, &bF);
    cL.getRgb(&rL, &gL, &bL);
    cR.getRgb(&rR, &gR, &bR);
    cRef.getRgb(&rRef, &gRef, &bRef);

    double red = getColorElem(intensity, specIntensity, alb, rF, rL, rR, rRef);
    double green = getColorElem(intensity, specIntensity, alb, gF, gL, gR, gRef);
    double blue = getColorElem(intensity, specIntensity, alb, bF, bL, bR, bRef);

    color.setRed(red);
    color.setGreen(green);
    color.setBlue(blue);

    return color;
}

std::tuple<bool, double, QVector3D, int> Drawing::sceneIntersect(const sight_t& sight)
{
    double t = std::numeric_limits<double>::max();
    int closest = -1;
    QVector3D norm;

    for (size_t i = 0; i < figures.size(); i++) {
        std::tuple<bool, double, QVector3D> res = figures[i]->rayIntersection(sight, lightSources);
        double resT = std::get<1>(res);
        if (std::get<0>(res) && resT < t) {
            t = resT;
            closest = i;
            norm = std::get<2>(res);
        }
    }

    if (closest == -1)
        return std::tuple<bool, double, QVector3D, int>(false, 0, QVector3D(0, 0, 0), 0);

    return std::tuple<bool, double, QVector3D, int>(true, t, norm, closest);
}

QColor Drawing::castRay(const sight_t& sight, const int& depth)
{
    double intensity = 0.0;
    double specIntensity = 0.0;

    std::tuple<bool, double, QVector3D, int> res = sceneIntersect(sight);
    if (depth > this->depth || !std::get<0>(res))
        return QColor(0, 0, 0);

    double t = std::get<1>(res);
    QVector3D norm = std::get<2>(res);
    int closest = std::get<3>(res);
    QVector3D cam = sight.cam;
    QVector3D dir = sight.dir;

    QVector3D intersection = cam + dir * t;

    QVector3D reflectDir = reflect(dir, norm).normalized();
    sight_t sightRefl = {
        .cam = intersection,
        .dir = reflectDir
    };
    QColor reflectCol = castRay(sightRefl, depth + 1);

    QVector3D refractDir = refract(dir, norm, figures[closest]->getMaterial().getRefCoef()).normalized();
    sight_t sightRefr = {
        .cam = intersection,
        .dir = refractDir
    };
    QColor refractCol = castRay(sightRefr, depth + 1);

    for (size_t k = 0; k < lightSources.size(); k++) {
        QVector3D light = intersection - lightSources[k].getPos();
        light = light.normalized();
        double t1 = 0.0;
        bool flag = false;
        sight_t sightLight = {
            .cam = lightSources[k].getPos(),
            .dir = light
        };
        std::tuple<bool, double, QVector3D> res1 = figures[closest]->rayIntersection(sightLight, lightSources);

        t1 = std::get<1>(res1);
        std::tuple<bool, double, QVector3D, int> res2 = sceneIntersect(sightLight);
        if (std::get<1>(res2) < t1)
            flag = true;

        if (!flag) {
            intensity += lightSources[k].getIntensity() * std::max(0.f, QVector3D::dotProduct(norm, (-1) * light));

            QVector3D mirrored = reflect(light, norm);
            mirrored = mirrored.normalized();
            double angle = QVector3D::dotProduct(mirrored, dir * (-1));
            specIntensity += powf(std::max(0.0, angle), figures[closest]->getMaterial().getSpecCoef()) * lightSources[k].getIntensity();
        }
    }

    QColor diffColor = figures[closest]->getMaterial().getDiffColor();
    QColor specColor = figures[closest]->getMaterial().getSpecColor();

    QColor color = getColor(intensity, specIntensity, figures[closest]->getMaterial().getAlbedo(), diffColor, specColor, reflectCol, refractCol);

    return color;
}

std::mutex mutex;

void Drawing::drawThread(QVector3D &preLim, QVector3D &lim, sight_t sight, std::vector<std::vector<uint>> &image)
{
    int startX = 0, endX = canvasWidth;

    int startY = preLim.y();
    if (preLim.x() != 0 || preLim.y() != 0) {
        startY--;
    }

    for (int y = startY; y < lim.y(); y++)
    {
        if (y == startY) {
            startX = preLim.x();
        } else {
            startX = 0;
        }

        if (y == lim.y() - 1) {
            endX = lim.x();
        } else {
            endX = canvasWidth;
        }

        for (int x = startX; x < endX; x++) {
            QVector3D pix = QVector3D(x, y, 200);
            QVector3D dir = (pix - sight.cam).normalized();

            sight.dir = dir;
            QColor refColor = castRay(sight, 0);
            image[y][x] = qRgb(refColor.red(), refColor.green(), refColor.blue());
        }
    }
}

std::shared_ptr<QImage> Drawing::drawFigures(int numThreads)
{
    std::shared_ptr<QImage> image = std::make_shared<QImage>(canvasWidth, canvasHeight, QImage::Format_RGB32);
    image->fill(Qt::black);

    QVector3D cam = QVector3D(0, 0, 3000);
    sight_t sight = {
        .cam = cam
    };

    int size = canvasHeight * canvasWidth;
    int cur = 0;
    int win = (size / numThreads);
    std::vector<QVector3D> limits;
    limits.push_back(QVector3D(0,0,200));

    std::vector<std::vector<uint>> colors(canvasHeight, std::vector<uint>(canvasWidth, 0));

    if (numThreads == 0) {
        limits.push_back(QVector3D(canvasWidth, canvasHeight, 200));
        drawThread(limits[0], limits[1], sight, colors);
    } else {
        for (int i = 0; i < numThreads; i++) {
            cur += win;
            limits.push_back(QVector3D(cur % canvasWidth, cur / canvasWidth, 200));
        }
        limits[limits.size() - 1] = QVector3D(canvasWidth, canvasHeight, 200);

        std::vector<std::thread> threads(numThreads);
        for (int i = 0; i < numThreads; i++) {
            threads[i] = std::thread(&Drawing::drawThread, this, std::ref(limits[i]), std::ref(limits[i+1]), sight, std::ref(colors));
        }

        for (int i = 0; i < numThreads; i++)
            threads[i].join();
    }

    for (int y = 0; y < canvasHeight; y++) {
        for (int x = 0; x < canvasWidth; x++) {
            image->setPixel(x, y, colors[y][x]);
        }
    }

    return image;
}

void Drawing::rotate(QVector3D angles)
{
    for (auto& figure : figures) {
        figure->rotate(angles);
    }
}
