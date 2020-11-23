#include<stdio.h>
#include<string.h>
int n;
int MyAtoi(const char* c){
	int a[1000],m=0,i,b;
	for(i=0;i<strlen(c);i++){
		if((c[i]<'0'||c[i]>'9')&&c[i]!='+'&&c[i]!='-') continue;
		b=i;
		a[0]=c[i];
		break;
	}
	int k=1;
	for(i=b+1;i<strlen(c);i++){
		if(c[i]<'0'||c[i]>'9') break;
		a[k]=c[i];
		k++;
	}
	if(a[0]!='-'&&a[0]!='+') m=a[0]-'0';
	for(i=1;i<k;i++){
		m=m*10+a[i]-'0';
	}
	if(a[0]=='-') m=-m;
	return m;
}
void main(){
	char str[1000];
	scanf("%s",str);
	n=MyAtoi(str);
	printf("%d\n",n);
}
