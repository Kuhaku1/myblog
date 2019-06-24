from django.contrib import admin
from django.urls import path
from . import views
app_name="mainsite"
urlpatterns = [
    path('body/', views.body),
]