from django.urls import path

from .views import (
    ActivateSegmentView,
    MoveSegmentView,
    PrintSegmentsView,
    RemoveSegmentView,
    SegmentCreateView,
    SegmentsListView,
)

urlpatterns = [
    path('add/', SegmentCreateView.as_view()),
    path('print/', PrintSegmentsView.as_view()),
    path('move/<int:pk>/', MoveSegmentView.as_view()),
    path('remove/<int:pk>/', RemoveSegmentView.as_view()),
    path('activate/<int:pk>/', ActivateSegmentView.as_view()),
    path('', SegmentsListView.as_view()),
]
