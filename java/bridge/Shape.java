package bridge;

abstract class Shape {
    Api Api;

    Shape(Api Api){
        this.Api = Api;
    }

    public abstract void Realdraw();
}
