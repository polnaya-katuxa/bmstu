#ifndef TEST_ROTATION_H
#define TEST_ROTATION_H

#include <QObject>

#include "rotation.h"

class TestRotation : public QObject {
    Q_OBJECT
public:
    explicit TestRotation(QObject* parent = nullptr);

private slots:
    void test_rotate_01();
};

#endif // TEST_ROTATION_H
