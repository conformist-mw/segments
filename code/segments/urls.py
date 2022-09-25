from django.urls import include, path

from .views import (
    ActivateSegmentView,
    CompaniesListView,
    MoveSegmentView,
    PrintSegmentsView,
    RemoveSegmentView,
    SectionsListView,
    SegmentCreateView,
    SegmentsListView,
)

segments_urls = [
    path('', SegmentsListView.as_view(), name='section_view'),
    path('add/', SegmentCreateView.as_view()),
    path('print/', PrintSegmentsView.as_view()),
    path('move/<int:pk>/', MoveSegmentView.as_view()),
    path('remove/<int:pk>/', RemoveSegmentView.as_view()),
    path('activate/<int:pk>/', ActivateSegmentView.as_view()),
]

urlpatterns = [
    path('', CompaniesListView.as_view(), name='home_view'),
    path('<slug:company>/', SectionsListView.as_view(), name='company_view'),
    path('<slug:company>/<slug:section>/', include(segments_urls)),
]
