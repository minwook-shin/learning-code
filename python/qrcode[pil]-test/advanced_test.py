import qrcode


qr = qrcode.QRCode(
    version=1,
    error_correction=qrcode.constants.ERROR_CORRECT_L,
    box_size=10,
    border=4,
)

qr.add_data('hello,world!')

img = qr.make_image(fill_color="white", back_color="black")
img.save("hello2.jpg")