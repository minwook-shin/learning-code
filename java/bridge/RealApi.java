package bridge;

class RealApi implements Api {
    public void draw(final double x, final double y, final double radius) {
        System.out.println("API1"+ " "+ x + " " + y+ " "+ radius);
    }
}