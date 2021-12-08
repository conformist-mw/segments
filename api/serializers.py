from rest_framework.serializers import (
    FloatField,
    IntegerField,
    ModelSerializer,
)
from rest_framework.exceptions import ValidationError

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


class SegmentDetailSerializer(ModelSerializer):
    color = ColorSerializer()
    rack = RackSerializer(read_only=False)
    racks = RackSerializer(many=True, read_only=True, source='get_racks')

    class Meta:
        model = Segment
        fields = '__all__'

    def update(self, instance, validated_data):
        validated_data.pop('rack')
        data = self.context['request'].data
        rack_id = data.get('rack', {}).get('id')
        if not rack_id:
            return super().update(instance, validated_data)
        try:
            rack = Rack.objects.get(id=rack_id)
        except Rack.DoesNotExist:
            raise ValidationError({'rack': 'Такого стеллажа не существует.'})
        if rack not in instance.get_racks():
            raise ValidationError({'rack': 'В данном цеху нет такого стеллажа'})
        instance.rack = rack
        instance.save()
        return super().update(instance, validated_data)


class SegmentSerializer(ModelSerializer):
    color = ColorSerializer()
    rack = RackSerializer()

    class Meta:
        model = Segment
        fields = '__all__'
