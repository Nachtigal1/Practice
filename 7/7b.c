#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include"../library/lib.h"
int main(){
float a=1,sum=0,x;
int k;
printf("Enter x:\n");
scanf("%f", &x);
if(x>0&&x<1){
for(k=1; a>=0.001; k++)
{
a=pow(-1,k)*((function18(k)*pow(x,k))/factorial(k));
sum+=a;
printf(" Sum:%f,k: %i", sum,k);
}    
}
else{
    printf("Incorrect x ");
}

return 0;
} 