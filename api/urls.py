from django.urls import include, path
from rest_framework_nested.routers import NestedDefaultRouter, SimpleRouter

from .views import CompanyViewSet, SectionViewSet, SegmentsViewSet

router = SimpleRouter()
router.register(r'companies', CompanyViewSet)

companies_router = NestedDefaultRouter(router, r'companies', lookup='company')
companies_router.register('sections', SectionViewSet)

sections_router = NestedDefaultRouter(
    companies_router, r'sections', lookup='section',
)
sections_router.register('segments', SegmentsViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('', include(companies_router.urls)),
    path('', include(sections_router.urls)),
    path('auth/', include('rest_framework.urls', namespace='rest_framework')),
]
