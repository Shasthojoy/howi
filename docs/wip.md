HOWI Community Edition is open source project governed by a [The MIT License][license-url].  
HOWI has been a proprietary [SDK][sdk-link], framework and package collection developed  
by Marko Kungla since 2005. It has modular `addon/plugin/lib` design style  
which as collection enables rapid development very wide range of software with  
clean and stable [API's][api-link]. Majority of HOWI was written in `C` while independent  
Addons, plugins and libraries cover [150+ programming languages][lang-stats-link] which could be  
used for very specific use cases within developed software as source or shared library.  

## Goals of 5th series of

- transform multitude of libraries from other language bindings to Go Lang.
- decide which Addons, Plugins and Libraries can be open sourced.
- decide which Addons, Plugins and Libraries will be deprecated.
- implement pre-release designs in Go for Addons, Plugins and Libraries which will be open sourced.

## Redesign with Go
Following describes HOWI-CE design for transforming it to Go Lang while keeping
HOWI's design style and philosophy.

- **HOWI Addons ./addon/<addon-name>**  
> Addons provide higher level API to include full feature set provided by that Addon and it's  
> Plugins to your software. Addons are safest to use since their high level API introduces rarely
> braking changes even when API's of it's plugins or used libraries change.

  - **Addons may import**
    - Go source
    - own plugins
    - any HOWI STD Library
    - any HOWI Library
    - external libraries
  - **Addons may never import**  
    - other Addons
    - Plugins from other Addons

- **HOWI Plugins ./addon/<addon-name>/plugin/<plugin-name>**  
> Plugins are always sub packages of Addon. While plugins are the core which Addons are made of they  
> are always usable as fully independent packages. Plugins are also safe to use and many cases more  
> reasonable to use when you only need features in your software which are provided that Plugin which  
> case importing/using entire Addon would bring unnecessary overhead. As noted earlier Plugins may
> introduce breaking changes earlier than Addons would so therefore that trade of should be taken  
> into consideration when planning your software.   

  - **Plugins may import**
    - Go source
    - any HOWI STD Library
    - any HOWI Library
  - **Plugins may never import**
    - other Addons

- **HOWI Libraries ./lib/<library-name>**  
> Libraries are low level and packages within HOWI. Purpose of these libraries is
> to provide low level bleeding edge features for HOWI Addons and their Plugins.
> HOWI Libraries mostly are extending or providing stable API for external libraries.
> (libraries can be nested)

  - **Libraries may import**
    - Go source
    - any HOWI STD Library
    - external libraries
    - any HOWI Library
  - **Libraries may never import**
    - any Addons
    - any Plugins

- **HOWI Standard Libraries ./std/<library-name>**  
> Standard Libraries are lowest level and most unstable packages within HOWI. Purpose of these
> libraries to provide low level bleeding edge features for HOWI Addons and their Plugins.
> HOWI Standard Libraries are often extending or replacing language features.
> (libraries can be nested)

  - **Libraries may import**
    - Go source
    - external libraries
  - **Libraries may never import**
    - any Addons
    - any Plugins
    - any HOWI Library

### As of june 2016 HOWI Version 4.x Language stats

- Final products had in total 87,960,753 lines of actual code  
- In total 17,921,779 lines of code comments  
- Left in total 12,145,167 blank lines into code  


1. 21.63% PHP 19,029,427 total lines
2. 20.31% C 17,861,652 total lines
3. 9.14% Java 8,043,654 total lines
4. 7.29% XML 6,412,496 total lines
5. 6.15% HTML 5,405,513 total lines
6. 4.97% C/C++ Header 4,374,468 total lines
7. 3.17% SQL 2,788,899 total lines
8. 2.69% Ruby 2,369,457 total lines
9. 2.33% CSS 2,051,853 total lines
10. 1.43% Python 1,255,578 total lines
11. 1.29% C++ 1,132,676 total lines
12. 0.84% JSON 734,936 total lines
13. 0.76% Mercury 667,025 total lines
14. 0.68% Pascal 602,294 total lines
15. 0.66% Racket 576,718 total lines
16. 0.58% Go 508,580 total lines
17. 0.57% Bourne Shell 505,756 total lines
18. 0.48% Kotlin 422,192 total lines
19. 0.47% Rust 411,739 total lines
20. 0.41% Scala 362,610 total lines
21. 0.41% Assembly 358,556 total lines
22. 0.35% YAML 304,504 total lines
23. 0.31% XAML 273,948 total lines
24. 0.29% VHDL 257,922 total lines
25. 0.29% Fortran 77 255,596 total lines
26. 0.28% D 249,087 total lines
27. 0.28% Perl 245,608 total lines
28. 0.24% Ada 208,225 total lines
29. 0.23% Lisp 198,745 total lines
30. 0.22% SAS 195,191 total lines
31. 0.22% C# 193,931 total lines
32. 0.22% Objective C 190,316 total lines
33. 0.21% Fortran 90 182,221 total lines
34. 0.20% yacc 180,054 total lines
35. 0.18% ASP.Net 161,069 total lines
36. 0.16% vim script 140,596 total lines
37. 0.16% Expect 139,086 total lines
38. 0.15% make 134,508 total lines
39. 0.15% OCaml 132,965 total lines
40. 0.14% m4 121,805 total lines
41. 0.14% ActionScript 120,721 total lines
42. 0.13% Hack 110,635 total lines
43. 0.12% Lua 101,694 total lines
44. 0.12% DTD 101,192 total lines
45. 0.11% Haskell 96,528 total lines
46. 0.11% Tcl/Tk 92,886 total lines
47. 0.10% TypeScript 87,853 total lines
48. 0.10% MATLAB 83,627 total lines
49. 0.10% SASS 83,570 total lines
50. 0.09% MSBuild script 80,851 total lines
51. 0.09% Groovy 74,888 total lines
52. 0.08% LESS 73,254 total lines
53. 0.08% lex 69,361 total lines
54. 0.08% Erlang 68,940 total lines
55. 0.07% Prolog 64,284 total lines
56. 0.07% Ioke 63,094 total lines
57. 0.07% HOWI-doc 62,010 total lines
58. 0.07% Grace 58,300 total lines
59. 0.06% Isabelle 55,679 total lines
60. 0.06% HyPhy 54,005 total lines
61. 0.06% Dart 53,637 total lines
62. 0.06% CoffeeScript 51,731 total lines
63. 0.06% Inform 7 49,827 total lines
64. 0.05% XSLT 46,673 total lines
65. 0.05% Verilog-SystemVerilog 45,787 total lines
66. 0.05% Standard ML 42,320 total lines
67. 0.05% Jupyter Notebook 39,847 total lines
68. 0.04% MXML 37,017 total lines
69. 0.04% XSD 36,794 total lines
70. 0.04% Idris 36,613 total lines
71. 0.04% IDL 36,249 total lines
72. 0.04% Bourne Again Shell 35,864 total lines
73. 0.04% Hy 34,075 total lines
74. 0.04% Objective C++ 32,524 total lines
75. 0.03% ERB 30,421 total lines
76. 0.03% Swift 29,414 total lines
77. 0.03% Julia 29,145 total lines
78. 0.03% JSONiq 28,495 total lines
79. 0.03% PowerShell 26,750 total lines
80. 0.03% Gosu 26,500 total lines
81. 0.03% Vala 24,836 total lines
82. 0.03% IGOR Pro 23,987 total lines
83. 0.03% Objective-J 23,984 total lines
84. 0.03% Jasmin 23,742 total lines
85. 0.03% R 22,908 total lines
86. 0.03% F# 22,495 total lines
87. 0.03% PureScript 22,086 total lines
88. 0.02% Visual Basic 20,639 total lines
89. 0.02% CUDA 18,880 total lines
90. 0.02% Maven 18,820 total lines
91. 0.02% JSP 18,568 total lines
92. 0.02% CMake 16,782 total lines
93. 0.02% Haml 14,621 total lines
94. 0.02% Windows Resource File 14,103 total lines
95. 0.01% Io 12,641 total lines
96. 0.01% Windows Module Definition 12,469 total lines
97. 0.01% Ant 12,325 total lines
98. 0.01% DOS Batch 12,049 total lines
99. 0.01% RobotFramework 10,675 total lines
100. 0.01% Inno Setup 9,620 total lines
101. 0.01% Protocol Buffers 9,590 total lines
102. 0.01% LOLCODE 9,283 total lines
103. 0.01% Cython 8,237 total lines
104. 0.01% Smarty 8,097 total lines
105. 0.01% Handlebars 8,041 total lines
106. 0.01% ASP 7,915 total lines
107. 0.01% Mustache 6,792 total lines
108. 0.01% Arduino Sketch 6,614 total lines
109. 0.01% LookML 6,482 total lines
110. 0.01% Vala Header 6,377 total lines
111. 0.01% ABAP 6,197 total lines
112. 0.01% Clojure 6,180 total lines
113. 0.01% LogTalk 4,754 total lines
114. sed 4,217 total lines
115. awk 3,782 total lines
116. Puppet 3,745 total lines
117. GLSL 3,400 total lines
118. NAnt script 3,372 total lines
119. J 3,372 total lines
120. SKILL 3,223 total lines
121. WiX source 3,101 total lines
122. JavaServer Faces 2,791 total lines
123. Pig Latin 2,775 total lines
124. QML 2,771 total lines
125. Visualforce Page 2,074 total lines
126. ColdFusion 2,061 total lines
127. xBase 1,893 total lines
128. Elixir 1,800 total lines
129. OpenCL 1,641 total lines
130. MUMPS 1,565 total lines
131. ColdFusion CFScript 1,551 total lines
132. dtrace 1,397 total lines
133. C Shell 1,078 total lines
134. Oracle Forms 903 total lines
135. Qt Project 867 total lines
136. Razor 764 total lines
137. diff 601 total lines
138. Ruby HTML 481 total lines
139. Korn Shell 441 total lines
140. WiX include 438 total lines
141. Windows Message File 346 total lines
142. HLSL 341 total lines
143. ClojureScript 189 total lines
144. Visualforce Component 173 total lines
145. Velocity Template Language 155 total lines
146. Javascript 123 total lines
147. Unity-Prefab 76 total lines
148. xBase Header 53 total lines
149. Grails 53 total lines
150. COBOL 52 total lines
151. Softbridge Basic 32 total lines
152. AutoHotkey 27 total lines
153. CCS 26 total lines
154. Fortran 95 3 total lines
155. Oracle Reports 3 total lines
