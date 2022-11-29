#include<stdio.h>
#include<math.h>
int main(){
    float A,ax=0,ay=0,B,bx=11,by=10,C,cx=-11,cy=12,AB,BC,AC,h,W,p;
    AB=(sqrt(pow((bx-ax),2)+pow((by-ay),2)));
    BC=(sqrt(pow((cx-bx),2)+pow((cy-by),2)));
    AC=(sqrt(pow((cx-ax),2)+pow((cy-ay),2)));
    p=((AB+BC+AC)/2);
    h=((2/AB)*sqrt(p*(p-AB)*(p-BC)*(p-AC)));
    W=(2*sqrt(BC*AC*p*(p-AB))/BC-AC);
    printf("h %f\n", h);
    printf("w %f\n", W);
    return 0;
}