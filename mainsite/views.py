from django.shortcuts import render

# Create your views here.
def body(request):
    return render(request,"mainsite/body.html")