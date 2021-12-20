from rest_framework.serializers import (
    CharField,
    FloatField,
    IntegerField,
    ModelSerializer,
)
from rest_framework.exceptions import ValidationError

from segments.models import Color, ColorType, Company, Rack, Section, Segment, OrderNumber


class OrderSerializer(ModelSerializer):

    class Meta:
        model = OrderNumber
        fields = ['name']


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


class SegmentDetailSerializer(ModelSerializer):
    color = ColorSerializer(read_only=True)
    racks = RackSerializer(many=True, read_only=True, source='get_racks')

    class Meta:
        model = Segment
        fields = [
            'width', 'height', 'square', 'created', 'deleted', 'defect',
            'description', 'active', 'color', 'order_number', 'racks', 'rack',
        ]
        read_only_fields = [
            'color', 'racks', 'width', 'height', 'square', 'created', 'deleted',
        ]

    def validate(self, attrs):
        if attrs.get('defect') and not attrs.get('description'):
            raise ValidationError('Описание обязательно при отметки дефекта.')
        return attrs


class SegmentListSerializer(ModelSerializer):
    color = ColorSerializer()
    rack = CharField(source='rack.name')
    order_number = OrderSerializer()

    class Meta:
        model = Segment
        fields = [
            'id', 'width', 'height', 'square', 'created', 'deleted', 'defect',
            'description', 'active', 'color', 'order_number', 'rack',
        ]
