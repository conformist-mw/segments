from rest_framework.serializers import (
    FloatField,
    IntegerField,
    ModelSerializer,
)

from segments.models import Color, ColorType, Company, Rack, Section, Segment


class CompanySerializer(ModelSerializer):

    class Meta:
        model = Company
        fields = ['name', 'slug', 'image']


class SectionSerializer(ModelSerializer):

    company = CompanySerializer()
    segments_count = IntegerField()
    racks_count = IntegerField()
    square_sum = FloatField()

    class Meta:
        model = Section
        fields = [
            'name',
            'slug',
            'company',
            'segments_count',
            'racks_count',
            'square_sum',
        ]


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
