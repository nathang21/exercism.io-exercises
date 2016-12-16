def hello(name=''):
	if name == '' or not isinstance(name, basestring):
		return "Hello, World!"
	
	return "Hello, " + name + "!"