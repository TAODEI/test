#include<stdio.h>
#include <stdlib.h>

typedef struct node{
	int Val;
	struct node* Prev;
	struct node* Next;
}listNode;

listNode* receive(listNode* ptr,int n){
	for (int i=0;i<n;i++){
		listNode* node=(listNode*)malloc(sizeof(listNode));
		scanf("%d", &node->Val);
		node->Next = ptr;
		node->Prev = NULL;
		if (ptr!=NULL) ptr->Prev = node;
		ptr = node;
	}
	return ptr;
}

listNode* creatlist(){
	listNode* head=(listNode*)malloc(sizeof(listNode));
	head->Next=NULL;
	return head;
}

void main(){
	int a,x=0;
	listNode* list=creatlist();
	listNode* ptr = NULL;
	scanf("%d",&a);
	ptr=receive(ptr,a);
	for(int i=0;i<a/2;i++){
		if(ptr[i]->Val!=ptr[a-i-1]->Val){
			x=1;
			break;
		}
	}
	if(x==0) puts("ture");
	else puts("false");
}
