from django.conf import settings
from django.contrib import admin
from django.urls import include, path

from segments.views import (
    ActivateSegmentView,
    MoveSegmentView,
    PrintSegmentsView,
    RemoveSegmentView,
    SegmentCreateView,
    SegmentsListView,
)

admin.site.site_header = 'Отрезки'

urlpatterns = [
    path('admin/', admin.site.urls),
    path('add/', SegmentCreateView.as_view()),
    path('print/', PrintSegmentsView.as_view()),
    path('move/<int:pk>/', MoveSegmentView.as_view()),
    path('remove/<int:pk>/', RemoveSegmentView.as_view()),
    path('activate/<int:pk>/', ActivateSegmentView.as_view()),
    path('', SegmentsListView.as_view()),
]

if settings.DEBUG:
    import debug_toolbar
    urlpatterns += [path('__debug__/', include(debug_toolbar.urls))]
