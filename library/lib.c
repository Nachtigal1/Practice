#include "./lib.h"
#include<math.h>
#include<stdlib.h>
float function13(float x){
    float f=( sin(pow(x,2))*cos(pow(x,3))-sin(x)+5.2);
    return f;
}
float function14(float x){
    float f=(2*sin(x)*sin(2*x-1.5)*cos(2*x+1.5)-6);
    return f;
}
float function15(float x){
    float f=(fabs(cos(pow(x,2))-0.51)*sin(3*x-4)*(-4.44));
    return f;
}
float function16(float x){
    float f=((cos(2.1*x)*abs(sin(x)))/0.15-5.8);
    return f;
 }
float function17(float x){
    float f=((fabs(cos(2*pow(x,3))+2*sin(x)/1.2))+10.51*fabs(cos(3*x)));
    return f;
 } 
float function18(float x){
    float f=(fabs(sin(pow(x,2)/1.5-2))+11.73*cos(1.6*x-1) );
    return f;
}
float factorial(float x){
    int fac;
    for(int k=1; k<=x; k++){
        fac*=k;
    }
    return fac;
}
float function19 (float x) {
    float f = 13.4*cos(fabs(x))*sin(pow(x,2)-2.25);
    return f;
}
float function20(float x){
    float f=(fabs(cos(pow(x,2)-3.8))/4.5-9.7*sin(x-3.1));
    return f;
}
float function21(float x){
    float f=13.4 * sin(-1.26) * cos(fabs(x / 7.5));
    return f;
}