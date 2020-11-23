#include<stdio.h>
void main(){
	int x;
	scanf("%d",&x);
	if(x>30){
		puts("fail");
		return;
	}
	float a=0;
	if(x<=10) a=x*0.8+0.2;
	if(x>10&&x<=20) a=0.2+x*0.75;
	if(x>20&&x<=30) a=0.2+x*0.7;
	printf("%.2f\n",a);
}
