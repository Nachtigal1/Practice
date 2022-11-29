#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include"../library/lib.h"
int main(){
    float x,y,a,b,c,d,e=2.71,i=11;
    printf("Enter x,a,b,c,d\n");
    scanf("%f", &x,&a,&b,&c,&d);
    if(fabs(x)<10){
    float fi=tan(x+a)-log(fabs(b+7))/log(i);
        y=function13(fi);
    }

    else if(fabs(x)>=10){
    float om=c*pow((pow(x,2)+d*pow(e,1.3)),1/5);
        y=function14(om);

    }
    printf("y=%f",y);
    return 0;

}