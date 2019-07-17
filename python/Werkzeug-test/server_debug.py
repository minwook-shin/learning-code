from werkzeug.debug import DebuggedApplication
import hello_world

app = DebuggedApplication(hello_world.application, evalex=True)