// Shows how to run three Steppers at once with varying speeds
//
// Requires the Adafruit_Motorshield v2 library
//   https://github.com/adafruit/Adafruit_Motor_Shield_V2_Library
// And AccelStepper with AFMotor support
//   https://github.com/adafruit/AccelStepper

// This tutorial is for Adafruit Motorshield v2 only!
// Will not work with v1 shields

#include <AccelStepper.h>
#include <Wire.h>
#include <Adafruit_MotorShield.h>
#include "utility/Adafruit_MS_PWMServoDriver.h"
#include <MultiStepper.h>

MultiStepper steppers;

Adafruit_MotorShield AFMStop(0x60); // Default address, no jumpers

// Connect two steppers with 200 steps per revolution (1.8 degree)
// to the top shield
Adafruit_StepperMotor *myStepper1 = AFMStop.getStepper(200, 1);
Adafruit_StepperMotor *myStepper2 = AFMStop.getStepper(200, 2);

//QUICK MATH
//Diameter of stepper shaft ~~ 0.2 inches (+0.03 for extra)
//Circumference of stepper shaft ~~ 0.725 inches of string per revolution

//500 steps = 3.674"
//1 step = .007348"


// you can change these to DOUBLE or INTERLEAVE or MICROSTEP!
// wrappers for the first motor!
void forwardstep1() {
  myStepper1->onestep(FORWARD, DOUBLE);
}
void backwardstep1() {
  myStepper1->onestep(BACKWARD, DOUBLE);
}
// wrappers for the second motor!
void forwardstep2() {
  myStepper2->onestep(FORWARD, DOUBLE);
}
void backwardstep2() {
  myStepper2->onestep(BACKWARD, DOUBLE);
}

// Now we'll wrap the 3 steppers in an AccelStepper object
AccelStepper stepper1(forwardstep1, backwardstep1);
AccelStepper stepper2(forwardstep2, backwardstep2);

void moveToPosition(int a, int b) {
  long positions[2]; // Array of desired stepper positions
  positions[0] = a;
  positions[1] = b;
  steppers.moveTo(positions);
  steppers.runSpeedToPosition(); // Blocks until all are in position
  
}

void moveToPosition2(int a, int b) {
  long positions[2]; //Array of desired stepper positions
  positions[0] = a;
  positions[1] = b;
  steppers.moveTo(positions);
}

void setup() {

  Serial.begin(9600);

  AFMStop.begin(); // Start the bottom shield

  stepper1.setMaxSpeed(2000.0);
  stepper1.setAcceleration(10.0);
  //  stepper1.moveTo(2000);

  stepper2.setMaxSpeed(2000.0);
  stepper2.setAcceleration(10.0);
  //  stepper2.moveTo(2000);

  steppers.addStepper(stepper1);
  steppers.addStepper(stepper2);

}


void loop() {


//  if (Serial.available() > 0) {
//    Serial.println("hello world");
    byte s;
    s = Serial.read();
    switch (s) {
      case 1:
        // motion
        unsigned char buf[4];
        
        Serial.readBytes(buf, 4);

//      Serial.println((unsigned int)buf[0]);
//      Serial.println((unsigned int)buf[1]);
//      Serial.println((unsigned int)buf[2]);
//      Serial.println((unsigned int)buf[3]);
        int a; 
        int b; 

        a = (buf[0] << 8) | buf[1];
        b = (buf[2] << 8) | buf[3];

//        Serial.write(1);
        
//      Serial.println(a);
//      Serial.println(b);

       moveToPosition(a, b); 

       Serial.write(1);
      break;
      case 2:
        stepper1.disableOutputs(); //Needs to be undone by calling enableOutputs() sometime later
        stepper2.disableOutputs();
        // disable
      break;
      case 3:
        stepper1.setCurrentPosition(stepper1.currentPosition()); //Sets the 'zero' position to the current position
        stepper2.setCurrentPosition(stepper2.currentPosition());
        // zero
      break;
    }
//  }

  //moveToPosition(0,-500);
  //stepper1.disableOutputs();
  //stepper2.disableOutputs();
//
//  moveToPosition(0, 0);
//
//  delay(3000);
//  moveToPosition(1700, 1700);
//  delay(1000);
//  moveToPosition(1700, 0);
//  delay(1000);
//  moveToPosition(0, 1700);
//  delay(1000);


}

