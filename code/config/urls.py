import debug_toolbar
from django.conf import settings
from django.contrib import admin
from django.urls import include, path

from segments.views import SegmentCreateView, SegmentsListView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('add/', SegmentCreateView.as_view()),
    path('', SegmentsListView.as_view()),
]

if settings.DEBUG:
    urlpatterns += [path('__debug__/', include(debug_toolbar.urls))]
