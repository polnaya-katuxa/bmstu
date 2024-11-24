#ifndef MATERIAL_H
#define MATERIAL_H

#include <QColor>
#include <QVector4D>

class Material {
public:
    Material();
    Material(QColor dC, QColor sC, QVector4D alb, double rCoef, double sCoef);
    Material(const Material& material);

    Material& operator=(const Material& material);
    Material& operator=(Material&& material);

    QColor getDiffColor();
    QColor getSpecColor();

    QVector4D getAlbedo();
    double getSpecCoef();
    double getRefCoef();

    void setDiffColor(QColor dC);
    void setSpecColor(QColor sC);

    void setAlbedo(QVector4D a);
    void setSpecCoef(double sC);
    void setRefCoef(double rC);

private:
    QColor diffColor;
    QColor specColor;

    QVector4D albedo;
    double refCoef;
    double specCoef;
};

#endif // MATERIAL_H
