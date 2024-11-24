#include "test_rotation.h"
#include <QtTest>

TestRotation::TestRotation(QObject* parent)
    : QObject { parent }
{
}

void TestRotation::test_rotate_01()
{
    double x = 1.0;
    double y = 2.0;

    rotate(x, y, 90.0);

    QCOMPARE(qFuzzyCompare(x, 2.0), true);
    QCOMPARE(qFuzzyCompare(y, -1.0), true);
}
