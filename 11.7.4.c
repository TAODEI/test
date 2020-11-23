#include<stdio.h>
typedef struct{
	int id;
	int sco;
}student;

student a[1000];
int n,i;

void getin(){
	scanf("%d%d",&a[i].id,&a[i].sco);
}
void sort(){
	int j;
	for(i=0;i<n;i++){
		for(j=i;j<n;j++){
			if(a[i].sco>a[j].sco){
				student tmp;
				tmp=a[i];
				a[i]=a[j];
				a[j]=tmp;
			}
		}
	}
}
void asc(){
	for(i=n-1;i>=0;i--){
		printf("%d %d\n",a[i].id,a[i].sco);
	}
}
void desc(){
	for(i=0;i<n;i++){
		printf("%d %d\n",a[i].id,a[i].sco);
	}
}
void main(){
	scanf("%d",&n);
	for(i=0;i<n;i++){
		getin();
	}
	sort();
	asc();
	desc();
}

