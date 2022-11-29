#include <stdio.h>
int main()
{
int m;
printf("Enter number of month: ");
scanf("%i",&m);
switch (m){
    case 1: case 2: case 3:
    printf("1 quarter");
    break;
    case 4: case 5: case 6:
    printf("2 quarter");
    break;
    case 7: case 8: case 9:
    printf("3 quarter");
    break;
    case 10: case 11: case 12:
    printf("4 quarter");
    break;
    case 0:
    printf("Error");
    break;
    default:
    break;
}
return 0;   
}