package bridge;

class CircleShape extends Shape {
    private double x, y, radius;
    
    public CircleShape(double x,double y,double radius,Api drawingAPI) {
        super(drawingAPI);
        this.x = x;  
        this.y = y;  
        this.radius = radius;
    }


    public void Realdraw() {
    	Api.draw(x, y, radius);
    }
}
