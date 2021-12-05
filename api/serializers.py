from rest_framework.serializers import ModelSerializer

from segments.models import Color, ColorType, Company, Rack, Section, Segment


class CompanySerializer(ModelSerializer):

    class Meta:
        model = Company
        fields = ['name', 'slug', 'image']


class SectionSerializer(ModelSerializer):

    company = CompanySerializer()

    class Meta:
        model = Section
        fields = ['name', 'slug', 'company']


class ColorTypeSerializer(ModelSerializer):

    class Meta:
        model = ColorType
        fields = ['name', 'slug']


class RackSerializer(ModelSerializer):

    class Meta:
        model = Rack
        fields = ['id', 'name']


class ColorSerializer(ModelSerializer):

    type = ColorTypeSerializer()

    class Meta:
        model = Color
        fields = ['id', 'name', 'slug', 'type']


class SegmentSerializer(ModelSerializer):
    color = ColorSerializer()
    rack = RackSerializer()

    class Meta:
        model = Segment
        fields = '__all__'
