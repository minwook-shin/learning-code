from geojson import Point

point = Point((126.97689056396484, 37.57880031671566))
print(point)
print(point.is_valid)

from geojson import MultiPoint


multi_p = MultiPoint([(126.97689056396484, 37.57880031671566), (126.97079658508302, 37.55492071859406)])
print(multi_p)
print(multi_p.is_valid)


from geojson import LineString


list_str = LineString([(126.95972442626953, 37.575671232633184), (126.99354171752928,37.57607938149231)])
print(list_str)
print(list_str.is_valid)


from geojson import MultiLineString


multi_str = MultiLineString([((126.95972442626953, 37.575671232633184), (126.99354171752928,37.57607938149231)),
                             [(126.95972442626953, 37.575671232633184), (126.99354171752928,37.57607938149231)]])
print(multi_str)
print(multi_str.is_valid)


from geojson import Polygon


polygon = Polygon([((126.98212623596191, 37.5774398615325), (126.96847915649414, 37.55839087912287),
                    (126.9938850402832, 37.55621354238461), (126.98212623596191, 37.5774398615325))])
print(polygon)
print(polygon.is_valid)


from geojson import MultiPolygon


multi_poly = MultiPolygon([([(126.98212623596191, 37.5774398615325), (126.96847915649414, 37.55839087912287),
                    (126.9938850402832, 37.55621354238461), (126.98212623596191, 37.5774398615325)],),
              ([(127.98212623596191, 37.5774398615325), (127.96847915649414, 37.55839087912287),
                    (127.9938850402832, 37.55621354238461), (127.98212623596191, 37.5774398615325)],)])
print(multi_poly)
print(multi_poly.is_valid)
