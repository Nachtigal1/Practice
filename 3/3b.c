#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include"../library/lib.h"
int main(){
    char A,B,Y,y;
    float i=11,salary,tax;
    printf("Enter type of job\n");
    scanf("%c",&y);
    if(y=='A'){
   salary=100*fabs(function13(i)+50);
    tax=salary*0.1;
    }
    else if(y=='B'){
    salary=150*fabs(function14(i)+100);
    tax=salary*0.15;
    }
    else if(y=='Y'){
    salary=200*fabs(function15(i)+135);
    tax=salary*0.15;
    }
    printf("Summa %f,summa tax %f, to issue %f",salary,tax,salary-tax);
    return 0;

}