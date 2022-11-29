#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include"../library/lib.h"
int main() {
float x=0,y=1,z;
for( int k=1; k<=16;k++)
{
 x+=function16(k);
 y*=function17(k);  
}
z=tan(x+y);
printf("x:%f, y:%f, z:%f", x,y,z);
return 0;  
}