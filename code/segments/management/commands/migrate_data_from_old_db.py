from collections import defaultdict

from django.core.management.base import BaseCommand

from segments.models import (
    Color,
    ColorType,
    OrderNumber,
    Rack,
    Segment,
    SegmentOld,
)


class Command(BaseCommand):
    help = 'Copy data from SegmentsOld to the new schema'

    def handle(self, *args, **options):
        add_colors_and_types()
        fill_segment_table()


def add_colors_and_types():

    racks = SegmentOld.objects.values_list('rack', flat=True).distinct()
    for rack in racks:
        Rack.objects.get_or_create(name=rack)

    types_with_colors = {
        'LAK': [
            'L 100', 'L 102', 'L 104', 'L 108', 'L 110', 'L 114', 'L 120',
            'L 140', 'L 156', 'L 160', 'L 162', 'L 180', 'L 201', 'L 205',
            'L 223', 'L 225', 'L 227', 'L 229', 'L 231', 'L 233', 'L 235',
            'L 303', 'L 303-3 Classic', 'L 305', 'L 307', 'L 309', 'L 311',
            'L 313', 'L 317', 'L 319', 'L 333', 'L 347', 'L 400', 'L 402',
            'L 404', 'L 406', 'L 408', 'L 410', 'L 412', 'L 416', 'L 420',
            'L 424', 'L 442', 'L 444', 'L 462', 'L 466', 'L 474', 'L 476',
            'L 478', 'L 490', 'L 501', 'L 507', 'L 511', 'L 515', 'L 519',
            'L 525', 'L 545', 'L 547', 'L 555', 'L 571', 'L 573', 'L 577',
            'L 604', 'L 606', 'L 608', 'L 610', 'L 614', 'L 618', 'L 624',
            'L 628', 'L 630', 'L 632', 'L 634', 'L 640', 'L 644', 'L 652',
            'L 664', 'L 666', 'L 674', 'L 682', 'L 684', 'L 707', 'L 713',
            'L 717', 'L 721', 'L 733', 'L 739', 'L 751', 'L 753', 'L 866',
        ],
        'SATIN': [
            'S 114', 'S 225', 'S 303', 'S 303-3 Classic', 'S 402', 'S 501',
            'S 507', 'S 511', 'S 652', 'S 717',
        ],
        'MAT': [
            'M 114', 'M 229', 'M 303', 'M 303-3 Classic', 'М 307', 'М 313',
            'M 319', 'М 347', 'M 501', 'M 507', 'M 511', 'M 652', 'M 717',
            'M 311', 'M 2 303 Обои', 'M 5 307 Обои', 'M 5 501 Обои',
            'M 5 507 Обои', 'M 7 Обои', 'M 8 Обои', 'M 9 Обои', 'M 10 Обои',
            'M 11 Обои', 'M 12 Обои', 'M 16 Обои', 'M 17 Обои',
        ],
        'EXCLUSIVE': [
            'Искра 100', 'Искра 114', 'Искра 160', 'Искра 180', 'Искра 225',
            'Искра 303', 'Искра 347', 'Искра 402', 'Искра 406', 'Искра 462',
            'Искра 478', 'Искра 511', 'Искра 571', 'Искра 652', 'Искра 733',
            'Нити', 'Вода', 'Облака', 'Cosmos', 'Металлик 901', 'Металлик 904',
            'Металлик 905', 'Металлик 906', 'Металлик 943', 'Металлик N320',
            'Металлик ST-81', 'Металлик ST-87', 'Металлик ST-94',
            'TRANSPARENT SATIN', 'TRANSPARENT LAK', 'DESCOR', 'Mat 10 gold',
            'Mat 11 gold', 'Mat 12 gold',
        ],
    }

    db_colors = defaultdict(set)
    for segment in SegmentOld.objects.all():
        db_colors[segment.type].add(segment.color)

    for color_type, colors in db_colors.items():
        if color_type not in types_with_colors:
            types_with_colors[color_type] = list(colors)
            continue
        default_colors = types_with_colors[color_type]
        different_colors = list(
            db_colors[color_type] - set(types_with_colors[color_type]),
        )
        types_with_colors[color_type] = default_colors + different_colors

    color_types_map = {}

    for color_type in types_with_colors:
        color_type_obj, _ = ColorType.objects.get_or_create(name=color_type)
        color_types_map[color_type] = color_type_obj

    for color_type, colors_values in types_with_colors.items():
        for color in colors_values:
            Color.objects.get_or_create(
                name=color, type=color_types_map[color_type],
            )


def fill_segment_table():

    racks = {}

    for segment in SegmentOld.objects.all():
        if segment.rack not in racks:
            rack, _ = Rack.objects.get_or_create(name=segment.rack)
            racks[segment.rack] = rack
        else:
            rack = racks[segment.rack]
        color, _ = Color.objects.get_or_create(
            name=segment.color, type__name=segment.type,
        )
        order_number = None
        if segment.order_number:
            order_number, _ = (
                OrderNumber.objects.get_or_create(name=segment.order_number)
            )
        defect = segment.defect if segment.defect is not None else False
        desc = segment.description if segment.description is not None else ''
        Segment.objects.create(
            width=segment.width,
            height=segment.height,
            square=segment.square,
            created=segment.created,
            deleted=segment.deleted,
            defect=defect,
            description=desc,
            active=segment.active,
            color=color,
            order_number=order_number,
            rack=rack,
        )
