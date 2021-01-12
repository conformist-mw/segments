import debug_toolbar
from django.conf import settings
from django.contrib import admin
from django.urls import include, path

from segments.views import (
    MoveSegmentView,
    PrintSegmentsView,
    SegmentCreateView,
    SegmentsListView,
)

urlpatterns = [
    path('admin/', admin.site.urls),
    path('add/', SegmentCreateView.as_view()),
    path('print/', PrintSegmentsView.as_view()),
    path('move/<int:pk>/', MoveSegmentView.as_view()),
    path('', SegmentsListView.as_view()),
]

if settings.DEBUG:
    urlpatterns += [path('__debug__/', include(debug_toolbar.urls))]
