#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include"../library/lib.h"
int main(){
float a=1,sum=0;
int k;
for(k=1; a>=0.001; k++)
{
a=(function18(k))/k;
sum+=a;
}
printf(" Sum:%f,k: %i", sum,k);
return 0;
} 