import qrcode.image.svg


factory = qrcode.image.svg.SvgImage
img = qrcode.make('hello,world!', image_factory=factory)
img.save("first.svg")

factory = qrcode.image.svg.SvgFragmentImage
img = qrcode.make('hello,world!', image_factory=factory)
img.save("second.svg")

factory = qrcode.image.svg.SvgPathImage
img = qrcode.make('hello,world!', image_factory=factory)
img.save("third.svg")
